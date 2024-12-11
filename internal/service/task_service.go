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
	ListTasks() []*domain.Task
	MarkTaskAsDone(taskID string) error
	GetTaskByID(taskID string) error
	DeleteTaskByID(taskID string) error
	UpdateTaskById(taskID string, newTask string) error
	ShowStatus() error
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

func (s *DefaultTaskService) ListTasks() []*domain.Task {
	tasks := s.repo.GetTasks()
	return tasks
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

	task.ToggleAsDone()

	if err := s.repo.SaveTasks(); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}

	status := "완료"
	if !task.Done {
		status = "미완료"
	}

	fmt.Printf("할일 %d번이 %s 상태로 변경되었습니다.\n", id, status)
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

	if err := s.repo.SaveTasks(); err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	fmt.Printf("Task %d has been deleted\n", id)
	return nil
}

func (s *DefaultTaskService) UpdateTaskById(taskID string, newTask string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", taskID)
	}

	if newTask == "" {
		return fmt.Errorf("new task cannot be empty")
	}

	if err := s.repo.UpdateTask(id, newTask); err != nil {
		return err
	}

	if err := s.repo.SaveTasks(); err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	fmt.Printf("Task %d has been updated\n", id)

	return nil
}

func (s *DefaultTaskService) ShowStatus() error {
	tasks := s.repo.GetTasks()

	total := len(tasks)
	completed := 0

	for _, task := range tasks {
		if task.Done {
			completed++
		}
	}

	fmt.Printf("전체 할일 수: %d\n", total)
	fmt.Printf("완료된 할일: %d\n", completed)
	fmt.Printf("미완료된 할일: %d\n", total-completed)

	return nil
}
