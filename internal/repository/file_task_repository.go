package repository

import (
	"encoding/json"
	"os"
	"slices"

	"todo-go-cli/internal/domain"
	apperrors "todo-go-cli/internal/errors"
)

type FileTaskRepository struct {
	tasks    []*domain.Task
	filePath string
}

func NewFileTaskRepository(filePath string) TaskRepository {
	repo := &FileTaskRepository{
		tasks:    make([]*domain.Task, 0),
		filePath: filePath,
	}

	if err := repo.LoadTasks(); err != nil {
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
		return nil, apperrors.NewNotFoundError("해당 번호의 할일을 찾을 수 없습니다")
	}

	return r.tasks[index], nil
}

func (r *FileTaskRepository) LoadTasks() error {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return apperrors.NewInternalError("할일 목록을 불러오는 중 오류가 발생했습니다", err)
	}

	if len(data) == 0 {
		r.tasks = make([]*domain.Task, 0)
		return nil
	}

	return json.Unmarshal(data, &r.tasks)
}

func (r *FileTaskRepository) SaveTasks() error {
	data, err := json.MarshalIndent(r.tasks, "", "    ")
	if err != nil {
		return apperrors.NewInternalError("할일 목록을 저장하는 중 오류가 발생했습니다", err)
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return apperrors.NewInternalError("할일 목록을 파일에 저장하는 중 오류가 발생했습니다", err)
	}

	return nil
}

func (r *FileTaskRepository) DeleteTasks(id int) error {
	oldLen := len(r.tasks)
	r.tasks = slices.DeleteFunc(r.tasks, func(t *domain.Task) bool {
		return t.ID == id
	})

	if oldLen == len(r.tasks) {
		return apperrors.NewNotFoundError("해당 번호의 할일을 찾을 수 없습니다")
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
