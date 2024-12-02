package main

import (
	"fmt"
	"os"

	"todo-go-cli/internal/config"
	apperrors "todo-go-cli/internal/errors"
	"todo-go-cli/internal/repository"
	"todo-go-cli/internal/service"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "설정을 초기화하는 중 오류가 발생했습니다:", err)
		os.Exit(1)
	}

	taskRepo := repository.NewFileTaskRepository(cfg.TasksFilePath)
	taskService := service.NewTaskService(taskRepo)

	if err := run(taskService, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "오류:", err)
		os.Exit(1)
	}
}

func run(taskService service.TaskService, args []string) error {
	if len(args) == 0 {
		printUsage()
		return nil
	}

	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			return apperrors.NewInvalidInputError("할일을 입력해주세요")
		}
		return taskService.AddTask(args[1])

	case "list":
		return taskService.ListTasks()

	case "done":
		if len(args) < 2 {
			return apperrors.NewInvalidInputError("할일 번호를 입력해주세요")
		}
		return taskService.ToggleTaskDone(args[1])

	case "get":
		if len(args) < 2 {
			return apperrors.NewInvalidInputError("할일 번호를 입력해주세요")
		}
		task, err := taskService.GetTask(args[1])
		if err != nil {
			return err
		}
		fmt.Println(task.String())
		return nil

	case "delete":
		if len(args) < 2 {
			return apperrors.NewInvalidInputError("할일 번호를 입력해주세요")
		}
		return taskService.DeleteTask(args[1])

	case "update":
		if len(args) < 3 {
			return apperrors.NewInvalidInputError("할일 번호와 새로운 내용을 입력해주세요")
		}
		return taskService.UpdateTask(args[1], args[2])

	case "status":
		return taskService.ShowStatus()

	default:
		printUsage()
		return nil
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add [task]        : 새로운 할일 추가")
	fmt.Println("  list              : 할일 목록 조회")
	fmt.Println("  done [taskID]     : 할일 완료 처리")
	fmt.Println("  get [taskID]      : 특정 할일 조회")
	fmt.Println("  delete [taskID]   : 할일 삭제")
	fmt.Println("  update [taskID] [new task] : 할일 수정")
	fmt.Println("  status            : 할일 통계 조회")
}
