package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "gocashflow/docs"

	"gocashflow/api/handler"
)

func NewServer() *fiber.App {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/customers", handler.GetAllCustomers)
	app.Get("/customers/:id", handler.GetCustomerByID)
	app.Post("/customers", handler.CreateCustomer)
	app.Delete("/customers/:id", handler.DeleteCustomerByID)
	app.Put("/customers/:id", handler.UpdateCustomer)
	app.Patch("/customers/:id", handler.PatchCustomer)
	app.Get("findcustomers/", handler.SearchCustomerByName)

	return app
}
