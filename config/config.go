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
	ID           int       `bson:"customer_id,omitempty" json:"id"`
	FirstName    string    `bson:"first_name" json:"first_name"`
	LastName     string    `bson:"last_name" json:"last_name"`
	Email        string    `bson:"email" json:"email" validate:"required,email"`
	PhoneNumber  string    `bson:"phone_number" json:"phone_number" validate:"required"`
	Address      string    `bson:"address" json:"address"`
	City         string    `bson:"city" json:"city"`
	State        string    `bson:"state" json:"state"`
	Country      string    `bson:"country" json:"country"`
	PostalCode   string    `bson:"postal_code" json:"postal_code"`
	DateOfBirth  time.Time `bson:"date_of_birth" json:"date_of_birth"`
	Gender       string    `bson:"gender" json:"gender"`
	Nationality  string    `bson:"nationality" json:"nationality"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
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




