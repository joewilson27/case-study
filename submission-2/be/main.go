package main

import (
	"fmt"
	"log"
	"os"

	"be/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setup App
	app := fiber.New()

	// CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH")
		c.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	// Seup Routes
	config.SetupRoutes(app)

	// check and load ENV
	config.LoadEnvir()

	// Setup Database
	config.SetupDatabase()

	fmt.Println("running on port: " + os.Getenv("APP_PORT"))

	// Run Apps on PORT
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
