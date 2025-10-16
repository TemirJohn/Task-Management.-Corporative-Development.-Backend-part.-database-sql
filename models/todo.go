package models

type Todo struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
