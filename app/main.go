package main

import (
	"boilerplate/app/routes"
	"boilerplate/config"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

var (
	port   = flag.String("port", ":3000", "Port to listen on")
	env    = flag.String("env", "dev", "Environment")
	engine = html.New("./app/public", ".html")
)

func main() {
	flag.Parse()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Congifure app
	config.New(*env)

	// Register routes
	routes.RegisterAPI(app)

	app.Static("/", "./app/public")

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}

		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Sorry, Not Found!",
		})
	})

	// Start server
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000

}
