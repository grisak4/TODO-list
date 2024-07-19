package models

type Task struct {
	TaskID    int    `json:"id"`
	TaskTitle string `json:"title"`
}
