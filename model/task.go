package model

type Task struct {
	Id       string `json:"id"`
	Task     string `json:"task"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
