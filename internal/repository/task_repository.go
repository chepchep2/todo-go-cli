package repository

import (
	"todo-go-cli/internal/domain"
)

// TaskRepository interface defines the contract for task persistence
type TaskRepository interface {
	GetTasks() []*domain.Task
	SaveTasks() error
	LoadTasks() error
	AddTask(task *domain.Task)
	FindTaskByID(id int) (*domain.Task, error)
	DeleteTasks(id int) error
	UpdateTask(id int, newTask string) error
}
