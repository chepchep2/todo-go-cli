package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

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
}

// FileTaskRepository implements TaskRepository using file storage
type FileTaskRepository struct {
	tasks    []*domain.Task
	filePath string
}

// NewFileTaskRepository creates a new FileTaskRepository instance
func NewFileTaskRepository(filePath string) TaskRepository {
	repo := &FileTaskRepository{
		tasks:    make([]*domain.Task, 0),
		filePath: filePath,
	}

	if err := repo.LoadTasks(); err != nil {
		// If file doesn't exist, start with empty tasks
		if os.IsNotExist(err) {
			repo.tasks = make([]*domain.Task, 0)
		}
	}

	return repo
}

func (r *FileTaskRepository) GetTasks() []*domain.Task {
	return r.tasks
}

func (r *FileTaskRepository) AddTask(task *domain.Task) {
	r.tasks = append(r.tasks, task)
}

func (r *FileTaskRepository) FindTaskByID(id int) (*domain.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	// task, found := slices.FindFunc(r.tasks, func(t *domain.Task) bool {
	// 	return t.ID == id
	// })

	// if !found {
	// 	return nil, fmt.Errorf("task with ID %d not found", id)
	// }

	return nil, fmt.Errorf("task with ID %d not found", id)
	// return task, nil
}

func (r *FileTaskRepository) LoadTasks() error {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return fmt.Errorf("failed to read tasks file: %w", err)
	}

	return json.Unmarshal(data, &r.tasks)
}

func (r *FileTaskRepository) SaveTasks() error {
	data, err := json.MarshalIndent(r.tasks, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	return os.WriteFile(r.filePath, data, 0644)
}

func (r *FileTaskRepository) DeleteTasks(id int) error {
	index := -1
	for i, task := range r.tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	r.tasks = slices.DeleteFunc(r.tasks, func(t *domain.Task) bool {
		return t.ID == id
	})

	return nil
}
