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
	app.Get("customers/", GetAllCustomers)
	app.Get("/customers/:id", GetCustomerByID)
	app.Post("/customers/", CreateCustomer)
	app.Delete("/customers/:id", DeleteCustomerByID)

	return app
}
