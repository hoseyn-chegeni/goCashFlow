package main

import (
	"fmt"
	"log"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "gocashflow/docs"
)


// @title Fiber Swagger API
// @version 1.0
// @description This is a sample Swagger + Fiber API
// @host localhost:3000
// @BasePath /

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



var client *mongo.Client

func ConnectToMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// آدرس کانکشن MongoDB
	mongoURI := "mongodb://localhost:27017"

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	// تست اتصال
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
}

func main() {
	ConnectToMongo()
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/hi/", SetupRoutes)

	log.Fatal(app.Listen(":3000"))
}
