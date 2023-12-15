package main

import (
	"GHB24Sync/app/routes"
	"GHB24Sync/config"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	env  = flag.String("env", "dev", "Environment")
)

func main() {
	flag.Parse()
	app := fiber.New()

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
