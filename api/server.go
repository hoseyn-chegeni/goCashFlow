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
	// CUSTOMERS
	app.Get("/customers", handler.GetAllCustomers)
	app.Get("/customers/:id", handler.GetCustomerByID)
	app.Post("/customers", handler.CreateCustomer)
	app.Delete("/customers/:id", handler.DeleteCustomerByID)
	app.Put("/customers/:id", handler.UpdateCustomer)
	app.Patch("/customers/:id", handler.PatchCustomer)
	app.Get("findcustomers/", handler.SearchCustomerByName)
	app.Patch("/customers/toggle-status/:id", handler.ToggleCustomerStatus)
	// ACCOUNTS
	app.Post("/accounts/", handler.CreateAccount)
	app.Get("/accounts/", handler.GetAllAccounts)
	app.Get("/accounts/:id", handler.GetAccountByID)
	app.Put("accounts/:id", handler.UpdateAccount)
	app.Delete("/accounts/:id", handler.DeleteAccount)


	return app
}
