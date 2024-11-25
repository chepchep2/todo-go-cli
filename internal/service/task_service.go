package service

import (
	"fmt"
	"strconv"

	"todo-go-cli/internal/domain"
	"todo-go-cli/internal/repository"
)

// TaskService defines the interface for task operations
type TaskService interface {
	AddTask(text string) error
	ListTasks()
	MarkTaskAsDone(taskID string) error
	GetTaskByID(taskID string) error
	DeleteTaskByID(taskID string) error
}

// DefaultTaskService implements TaskService
type DefaultTaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new DefaultTaskService instance
func NewTaskService(repo repository.TaskRepository) TaskService {
	return &DefaultTaskService{
		repo: repo,
	}
}

func (s *DefaultTaskService) AddTask(text string) error {
	if text == "" {
		return fmt.Errorf("task text cannot be empty")
	}

	tasks := s.repo.GetTasks()
	newTask := domain.NewTask(len(tasks)+1, text)
	s.repo.AddTask(newTask)

	if err := s.repo.SaveTasks(); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}

	fmt.Printf("Task added: %s\n", text)
	return nil
}

func (s *DefaultTaskService) ListTasks() {
	tasks := s.repo.GetTasks()
	if len(tasks) == 0 {
		fmt.Println("할일이 없습니다")
		return
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}
}

func (s *DefaultTaskService) MarkTaskAsDone(taskID string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", taskID)
	}

	task, err := s.repo.FindTaskByID(id)
	if err != nil {
		return err
	}

	task.MarkAsDone()

	if err := s.repo.SaveTasks(); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}

	fmt.Printf("할일 %d번이 완료되었습니다.\n", id)
	return nil
}

func (s *DefaultTaskService) GetTaskByID(taskID string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", taskID)
	}

	task, err := s.repo.FindTaskByID(id)
	if err != nil {
		return err
	}

	fmt.Println(task.String())
	return nil
}

func (s *DefaultTaskService) DeleteTaskByID(taskID string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", taskID)
	}

	err = s.repo.DeleteTasks(id)
	if err != nil {
		return err
	}

	fmt.Printf("Task %d has been deleted\n", id)
	return nil
}
