package routes

import "github.com/gofiber/fiber/v2"

func RegisterAPI(app fiber.Router) {
	api := app.Group("/api")
	api.Group("/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
