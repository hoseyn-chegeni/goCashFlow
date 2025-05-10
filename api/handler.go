package api

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"
	"gocashflow/config"
)

// SetupRoutes handles GET /hi
// @Summary SetupRoutes Endpoint
// @Description Responds with a simple "hi!" message.
// @Tags SetupRoutes
// @Accept json
// @Produce json
// @Success 200 {string} string "SetupRoutes!"
// @Router /hi/ [get]
func SetupRoutes(c *fiber.Ctx) error {
	return c.SendString("hi")
}



// SetupRoutes handles POST /customers/create
// @Summary CreateCustomers Endpoint
// @Description Responds with a simple "Create Customers!" message.
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {string} string "Customer add successfully!"
// @Router /customers/create/ [post]
func CreateCustomer(c *fiber.Ctx) error {
	var customer config.Customer

	// Parse the request body into the Customer struct
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Set the timestamps
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	// Insert the customer into the MongoDB collection
	_, err := config.CustomerCollection.InsertOne(context.Background(), customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert customer"})
	}

	// Return a success response with the created customer details
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Customer created successfully!",
		"customer": customer,
	})
}
