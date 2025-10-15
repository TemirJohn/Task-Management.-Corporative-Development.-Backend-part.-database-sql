package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // драйвер PostgreSQL
)

var DB *sql.DB

func ConnectDB() {
	connStr := "host=localhost user=temirzhan password=11466795 dbname=todolist_sql port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to PostgreSQL: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Database not reachable: %v", err))
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		description TEXT
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Error creating table: %v", err))
	}

	DB = db
	//fmt.Println("Connected to PostgreSQL using database/sql!")
}
