package config

import (
	"be/config/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// api
	api := app.Group("/api")

	// checker health
	checkerRoutes := api.Group("/checker-health")
	routes.CheckerHealth(checkerRoutes)

	// task
	taskRoutes := api.Group("/task")
	routes.Task(taskRoutes)

}
