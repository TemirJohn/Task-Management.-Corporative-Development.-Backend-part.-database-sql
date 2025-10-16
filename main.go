package main

import (
	"fmt"
	"time"

	"todo_list_sql/config"
	"todo_list_sql/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	config.ConnectDB()

	app.GET("/api/todos", handlers.FetchAllTodos)
	app.GET("/api/todos/:id", handlers.GetTodoByID)
	app.POST("/api/todos", handlers.CreateTodo)
	app.PUT("/api/todos/:id", handlers.UpdateTodo)
	app.DELETE("/api/todos/:id", handlers.DeleteTodo)

	app.Run(fmt.Sprintf(":%v", 8080))
}
