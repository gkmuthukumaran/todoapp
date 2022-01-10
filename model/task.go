package model

type Task struct {
	Id       string `json:"id"`
	Task     string `json:"task"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type Category struct {
	Id       string `json:"id"`
	Category string `json:"category"`
}
