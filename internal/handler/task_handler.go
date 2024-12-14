package handler

import (
	"todo-go-cli/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// POST /todos/add
func (h *TaskHandler) AddTask(c *fiber.Ctx) error {
	var request struct {
		Text string `json:"text"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.taskService.AddTask(request.Text); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
	})
}

// GET /todos/list
func (h *TaskHandler) ListTasks(c *fiber.Ctx) error {
	tasks := h.taskService.ListTasks()
	return c.JSON(tasks)
}

// GET /todos/id
func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	if err := h.taskService.GetTaskByID(taskID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

// PUT /todos/id
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	var request struct {
		Text string `json:"text"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.taskService.UpdateTaskById(taskID, request.Text); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
	})
}

// PUT /todos/id/toggle
func (h *TaskHandler) ToggleTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	if err := h.taskService.MarkTaskAsDone(taskID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task toggled successfully",
	})
}

// GET /todos/status
func (h *TaskHandler) GetStatus(c *fiber.Ctx) error {
	if err := h.taskService.ShowStatus(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

// DELETE /todos/id
func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	if err := h.taskService.DeleteTaskByID(taskID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
