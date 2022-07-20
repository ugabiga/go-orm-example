package gorme

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Birthday  sql.NullTime `gorm:""`
	Tasks     []Task       `gorm:"foreignKey:UserID"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
}

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User `gorm:"association_foreignkey:ID"`
	Title     string
	Note      string
	Status    string     `gorm:"default:'todo'"`
	Projects  []*Project `gorm:"many2many:project_tasks;"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

type ProjectStatus string

const (
	ProjectStatusTodo       TaskStatus = "todo"
	ProjectStatusInProgress TaskStatus = "in_progress"
	ProjectStatusDone       TaskStatus = "done"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string  `gorm:"default:'todo'"`
	Tasks       []*Task `gorm:"many2many:project_tasks;"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
