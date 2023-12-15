package routes

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app fiber.Router) {
	env := os.Getenv("APP_ENV")

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World 👋!\n" + "Success: True\n" + "Env: " + env + "\n" + "Timestamp: " + time.Now().String() + "\n")
		return c.Render("index", fiber.Map{
			"Message":   "Hello, World 👋!",
			"Status":    "success",
			"Env":       env,
			"Timestamp": time.Now().String(),
		})
	})

	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "Hello, World 👋!",
			"status":    "success",
			"env":       env,
			"timestamp": time.Now().String(),
		})
	})
}
