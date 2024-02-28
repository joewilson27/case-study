package routes

import (
	"be/pkg/core/task"

	"github.com/gofiber/fiber/v2"
)

func Task(taskRoutes fiber.Router) {

	taskRoutes.Get("/", task.GetTasks)

	taskRoutes.Get("/item/:id", task.GetTaskById)

	taskRoutes.Post("/", task.AddTask)

	taskRoutes.Delete("/:id", task.DeleteTask)

	taskRoutes.Patch("/:id", task.UpdateTask)

	taskRoutes.Get("/complete", task.GetTaskComplete)
}
