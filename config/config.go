package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Customer struct {
	ID           int       `bson:"customer_id,omitempty"`  // Customer ID (Auto Increment not supported natively in MongoDB)
	FirstName    string    `bson:"first_name"`
	LastName     string    `bson:"last_name"`
	Email        string    `bson:"email"`
	PhoneNumber  string    `bson:"phone_number"`
	Address      string    `bson:"address"`
	City         string    `bson:"city"`
	State        string    `bson:"state"`
	Country      string    `bson:"country"`
	PostalCode   string    `bson:"postal_code"`
	DateOfBirth  time.Time `bson:"date_of_birth"`
	Gender       string    `bson:"gender"`
	Nationality  string    `bson:"nationality"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}


var Client *mongo.Client
var CustomerCollection *mongo.Collection

func ConnectToMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "mongodb://localhost:27017"

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
	CustomerCollection = Client.Database("cashflow").Collection("customers")
}




