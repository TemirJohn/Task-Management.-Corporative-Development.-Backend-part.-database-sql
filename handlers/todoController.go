package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"todo_list_sql/config"
	"todo_list_sql/models"

	"github.com/gin-gonic/gin"
)

// GET /api/todos
func FetchAllTodos(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, description, completed FROM todos ORDER BY id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, todos)
}

// GET /api/todos/:id
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	var t models.Todo
	err := config.DB.QueryRow("SELECT id, title, description, completed FROM todos WHERE id=$1", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Completed)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

// POST /api/todos
func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, input.Title, input.Description, input.Completed).Scan(&input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// PUT /api/todos/:id
func UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE todos
		SET title=$1, description=$2, completed=$3
		WHERE id=$4
		RETURNING id, title, description, completed
	`
	err = config.DB.QueryRow(query, input.Title, input.Description, input.Completed, id).
		Scan(&input.ID, &input.Title, &input.Description, &input.Completed)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

// DELETE /api/todos/:id
func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	_, err := config.DB.Exec("DELETE FROM todos WHERE id=$1", idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
