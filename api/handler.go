package api

import (
	"fmt"
	"time"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gocashflow/config"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)



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



// @Summary Delete a customer
// @Tags Customers
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /customers/{id} [delete]
func DeleteCustomerByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := config.CustomerCollection.DeleteOne(ctx, bson.M{"customer_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete customer",
		})
	}

	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Customer with ID %d deleted successfully", id),
	})
}



// @Summary Full update of customer
// @Tags Customers
// @Param id path int true "Customer ID"
// @Accept json
// @Produce json
// @Success 200 {object} config.Customer
// @Router /customers/{id} [put]
func UpdateCustomer(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var customer config.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	customer.UpdatedAt = time.Now()
	customer.CreatedAt = time.Now() // Optional
	customer.ID = id

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"customer_id": id}
	update := bson.M{"$set": customer}

	res, err := config.CustomerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update failed"})
	}

	if res.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.JSON(fiber.Map{"message": "Customer updated successfully", "customer": customer})
}


// @Summary Partial update of customer
// @Tags Customers
// @Param id path int true "Customer ID"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /customers/{id} [patch]
func PatchCustomer(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	updateData["updated_at"] = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"customer_id": id}
	update := bson.M{"$set": updateData}

	res, err := config.CustomerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Patch update failed"})
	}

	if res.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.JSON(fiber.Map{"message": "Customer updated successfully"})
}
