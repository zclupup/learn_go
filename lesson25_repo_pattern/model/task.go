package model

import "time"

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type CreateTaskRequest struct {
	Title string
}
