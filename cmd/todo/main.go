package main

import (
	"fmt"
	"os"

	"todo-go-cli/internal/config"
	"todo-go-cli/internal/repository"
	"todo-go-cli/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error initializing config:", err)
		os.Exit(1)
	}

	repo := repository.NewFileTaskRepository(cfg.TasksFilePath)
	taskService := service.NewTaskService(repo)

	if len(os.Args) > 1 && os.Args[1] == "server" {
		runServer(taskService)
	} else {
		if err := run(taskService, os.Args); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}
}

func runServer(taskService service.TaskService) {
	app := fiber.New()

	app.Post("/todos/add", func(c *fiber.Ctx) error {
		var input struct {
			Text string `json:"text"`
		}
		if err := c.BodyParser(&input); err != nil {
			return err
		}

		if err := taskService.AddTask(input.Text); err != nil {
			return err
		}

		return c.JSON(fiber.Map{"message": "Task added successfully"})
	})

	app.Get("/todos/list", func(c *fiber.Ctx) error {
		tasks := taskService.ListTasks()
		return c.JSON(tasks)
	})

	app.Listen(":8080")
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

	case "delete":
		if len(args) < 3 {
			return fmt.Errorf("list 번호를 입력하세요")
		}

		return taskService.DeleteTaskByID(args[2])

	case "update":
		if len(args) < 4 {
			return fmt.Errorf("update 번호와 update 내용을 입력하세요")
		}
		return taskService.UpdateTaskById(args[2], args[3])

	case "status":
		return taskService.ShowStatus()
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
	fmt.Println("  todo delete [id]   Delete a task by ID")
	fmt.Println("  todo status        show task's status")
}
