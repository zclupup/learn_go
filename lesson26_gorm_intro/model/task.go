package model

import "time"

type Task struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"column:title;not null" json:"title"`
	Done      bool      `gorm:"column:done;not null" json:"done"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (Task) TableName() string {
	return "lesson26_tasks"
}
