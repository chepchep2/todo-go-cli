package main

import (
	"log"
	"todo-go-cli/internal/handler"
	"todo-go-cli/internal/repository"
	"todo-go-cli/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 저장소 초기화 - data 디렉토리에 저장
	repo := repository.NewFileTaskRepository("data/tasks.json")

	// 서비스 초기화
	taskService := service.NewTaskService(repo)

	// 핸들러 초기화
	taskHandler := handler.NewTaskHandler(taskService)

	// Fiber 앱 생성
	app := fiber.New()

	// 라우트 설정
	todos := app.Group("/todos")
	todos.Post("/add", taskHandler.AddTask)
	todos.Get("/list", taskHandler.ListTasks)
	todos.Get("/status", taskHandler.GetStatus)
	todos.Get("/:id", taskHandler.GetTask)
	todos.Put("/:id", taskHandler.UpdateTask)
	todos.Put("/:id/toggle", taskHandler.ToggleTask)
	todos.Delete("/:id", taskHandler.DeleteTask)

	// 서버 시작
	log.Fatal(app.Listen(":8080"))
}
