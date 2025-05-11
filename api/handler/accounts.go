package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"gocashflow/config"
)

// ✅ Create Account
func CreateAccount(c *fiber.Ctx) error {
	var account config.Account
	if err := c.BodyParser(&account); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	count, _ := config.AccountCollection.CountDocuments(context.Background(), bson.M{})
	account.ID = int(count) + 1
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	account.Status = "Active"

	_, err := config.AccountCollection.InsertOne(context.Background(), account)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create account"})
	}
	return c.Status(201).JSON(account)
}

// ✅ Get all accounts
func GetAllAccounts(c *fiber.Ctx) error {
	cursor, err := config.AccountCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch accounts"})
	}
	var accounts []config.Account
	if err := cursor.All(context.Background(), &accounts); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Decode error"})
	}
	return c.JSON(accounts)
}

// ✅ Get account by ID
func GetAccountByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var account config.Account
	err := config.AccountCollection.FindOne(context.Background(), bson.M{"account_id": id}).Decode(&account)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Account not found"})
	}
	return c.JSON(account)
}

// ✅ Update account (PUT)
func UpdateAccount(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated config.Account
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	updated.UpdatedAt = time.Now()

	res, err := config.AccountCollection.UpdateOne(context.Background(),
		bson.M{"account_id": id},
		bson.M{"$set": updated},
	)
	if err != nil || res.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Account not found"})
	}
	return c.JSON(fiber.Map{"message": "Account updated"})
}

// ✅ Delete account
func DeleteAccount(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res, err := config.AccountCollection.DeleteOne(context.Background(), bson.M{"account_id": id})
	if err != nil || res.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Account not found"})
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Account %d deleted", id)})
}