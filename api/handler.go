package api

import (
	"time"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gocashflow/config"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
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

	// اعتبارسنجی
	if err := validate.Struct(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// 🆕 گرفتن تعداد رکوردهای موجود برای ساخت ID جدید
	count, err := config.CustomerCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to count documents"})
	}
	customer.ID = int(count) + 1

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	_, err = config.CustomerCollection.InsertOne(context.Background(), customer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert customer"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Customer created successfully!",
		"customer": customer,
	})
}



// @Summary Get all customers
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {array} config.Customer
// @Router /customers [get]
func GetAllCustomers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.CustomerCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch customers",
		})
	}

	var customers []config.Customer
	if err := cursor.All(ctx, &customers); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode customers",
		})
	}

	return c.JSON(customers)
}

// @Summary Get customer by ID
// @Tags Customers
// @Param id path int true "Customer ID"
// @Accept json
// @Produce json
// @Success 200 {object} config.Customer
// @Router /customers/{id} [get]
func GetCustomerByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid customer ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var customer config.Customer
	err = config.CustomerCollection.FindOne(ctx, bson.M{"customer_id": id}).Decode(&customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.JSON(customer)
}