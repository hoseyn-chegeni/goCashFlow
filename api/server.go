package api

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "gocashflow/docs" // required for Swagger docs
)

// NewServer initializes and returns a configured Fiber app
func NewServer() *fiber.App {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/hi/", SetupRoutes)

	return app
}
