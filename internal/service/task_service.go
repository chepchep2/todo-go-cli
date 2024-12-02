package service

import (
	"fmt"
	"strconv"
	"todo-go-cli/internal/domain"
	apperrors "todo-go-cli/internal/errors"
	"todo-go-cli/internal/repository"
)

// TaskService defines the interface for task operations
type TaskService interface {
	AddTask(text string) error
	ListTasks() error
	ToggleTaskDone(taskID string) error
	GetTask(taskID string) (*domain.Task, error)
	DeleteTask(taskID string) error
	UpdateTask(taskID string, newText string) error
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
		return apperrors.NewInvalidInputError("할일 내용을 입력해주세요")
	}

	tasks := s.repo.GetTasks()
	newTask := domain.NewTask(len(tasks)+1, text)
	s.repo.AddTask(newTask)

	if err := s.repo.SaveTasks(); err != nil {
		return apperrors.NewInternalError("할일을 저장하는 중 오류가 발생했습니다", err)
	}

	fmt.Printf("할일이 추가되었습니다: %s\n", text)
	return nil
}

func (s *DefaultTaskService) ListTasks() error {
	tasks := s.repo.GetTasks()
	if len(tasks) == 0 {
		fmt.Println("할일이 없습니다")
		return nil
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}
	return nil
}

func (s *DefaultTaskService) ToggleTaskDone(taskID string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return apperrors.NewInvalidInputError("유효하지 않은 할일 번호입니다: " + taskID)
	}

	task, err := s.repo.FindTaskByID(id)
	if err != nil {
		return apperrors.NewNotFoundError("해당 번호의 할일을 찾을 수 없습니다: " + taskID)
	}

	task.ToggleAsDone()

	if err := s.repo.SaveTasks(); err != nil {
		return apperrors.NewInternalError("할일을 저장하는 중 오류가 발생했습니다", err)
	}

	status := "완료"
	if !task.Done {
		status = "미완료"
	}

	fmt.Printf("할일 %d번이 %s 상태로 변경되었습니다.\n", id, status)
	return nil
}

func (s *DefaultTaskService) GetTask(taskID string) (*domain.Task, error) {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return nil, apperrors.NewInvalidInputError("유효하지 않은 할일 번호입니다: " + taskID)
	}

	return s.repo.FindTaskByID(id)
}

func (s *DefaultTaskService) DeleteTask(taskID string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return apperrors.NewInvalidInputError("유효하지 않은 할일 번호입니다: " + taskID)
	}

	err = s.repo.DeleteTasks(id)
	if err != nil {
		return apperrors.NewNotFoundError("해당 번호의 할일을 찾을 수 없습니다: " + taskID)
	}

	if err := s.repo.SaveTasks(); err != nil {
		return apperrors.NewInternalError("할일을 저장하는 중 오류가 발생했습니다", err)
	}

	fmt.Printf("할일 %d번이 삭제되었습니다\n", id)
	return nil
}

func (s *DefaultTaskService) UpdateTask(taskID string, newText string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return apperrors.NewInvalidInputError("유효하지 않은 할일 번호입니다: " + taskID)
	}

	if newText == "" {
		return apperrors.NewInvalidInputError("새로운 할일 내용을 입력해주세요")
	}

	if err := s.repo.UpdateTask(id, newText); err != nil {
		return apperrors.NewNotFoundError("해당 번호의 할일을 찾을 수 없습니다: " + taskID)
	}

	if err := s.repo.SaveTasks(); err != nil {
		return apperrors.NewInternalError("할일을 저장하는 중 오류가 발생했습니다", err)
	}

	fmt.Printf("할일 %d번이 수정되었습니다\n", id)

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
