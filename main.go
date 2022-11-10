package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173",
	}))

	app.Static("/", "./client/dist")

	app.Get("/api/helloWorld", func(c *fiber.Ctx) error {
		return c.SendString("Hello from the server!")
	})

	app.Listen(":3000")
}
