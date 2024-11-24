package main

import (
	"fmt"
	"os"

	"todo-go-cli/internal/config"
	"todo-go-cli/internal/repository"
	"todo-go-cli/internal/service"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error initializing config:", err)
		os.Exit(1)
	}

	repo := repository.NewFileTaskRepository(cfg.TasksFilePath)
	taskService := service.NewTaskService(repo)

	if err := run(taskService, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(taskService service.TaskService, args []string) error {
	if len(args) < 2 {
		printUsage()
		return nil
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			return fmt.Errorf("할일을 입력하세요")
		}
		return taskService.AddTask(args[2])

	case "list":
		taskService.ListTasks()
		return nil

	case "done":
		if len(args) < 3 {
			return fmt.Errorf("list 번호를 입력하세요")
		}
		return taskService.MarkTaskAsDone(args[2])

	case "get":
		if len(args) < 3 {
			return fmt.Errorf("list 번호를 입력하세요")
		}

		return taskService.GetTaskByID(args[2])
	default:
		printUsage()
		return nil
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo add [task]    Add a new task")
	fmt.Println("  todo list          Show all tasks")
	fmt.Println("  todo done [id]     Mark a task as done")
	fmt.Println("  todo get [id]      Get a task by ID")
}
