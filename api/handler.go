package api

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"
	"gocashflow/config"
	"github.com/go-playground/validator/v10"
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


var validate = validator.New()
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

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// ✅ اعتبارسنجی انجام می‌دیم
	if err := validate.Struct(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
			"details": err.Error(),
		})
	}

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	_, err := config.CustomerCollection.InsertOne(context.Background(), customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert customer"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Customer created successfully!",
		"customer": customer,
	})
}
