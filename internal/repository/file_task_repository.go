package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"todo-go-cli/internal/domain"
)

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
	index := slices.IndexFunc(r.tasks, func(t *domain.Task) bool {
		return t.ID == id
	})

	if index == -1 {
		return nil, fmt.Errorf("task with ID %d not found", id)
	}

	return r.tasks[index], nil
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
	oldLen := len(r.tasks)

	r.tasks = slices.DeleteFunc(r.tasks, func(t *domain.Task) bool {
		return t.ID == id
	})

	if oldLen == len(r.tasks) {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return nil
}

func (r *FileTaskRepository) UpdateTask(id int, newTask string) error {
	task, err := r.FindTaskByID(id)
	if err != nil {
		return err
	}

	task.Text = newTask

	return nil
}
