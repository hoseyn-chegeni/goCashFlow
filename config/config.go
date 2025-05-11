package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI        string
	Database   string
	Collection string
}


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
	Status       bool      `bson:"status" json:"status"`
}

type Account struct {
	ID          int       `bson:"account_id,omitempty" json:"id"`
	CustomerID  int       `bson:"customer_id" json:"customer_id"`               // foreign key
	AccountType string    `bson:"account_type" json:"account_type"`             // "Checking", "Savings", etc.
	Balance     float64   `bson:"balance" json:"balance"`                       // consider using float64 for simplicity
	Status      string    `bson:"status" json:"status"`                         // "Active", "Suspended", "Closed"
	IsPrimary   bool      `bson:"is_primary" json:"is_primary"`                 // optional: is this the customer's main account?
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

var (
	Client              *mongo.Client
	CustomerCollection  *mongo.Collection
	MongoSettings       MongoConfig
	AccountCollection *mongo.Collection
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("❌ Error reading config file: %v", err)
	}

	MongoSettings = MongoConfig{
		URI:        viper.GetString("mongodb.uri"),
		Database:   viper.GetString("mongodb.database"),
		Collection: viper.GetString("mongodb.collection"),
	}
}

func ConnectToMongo() {
	LoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(MongoSettings.URI))
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	if err := Client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	db := Client.Database(MongoSettings.Database)
	CustomerCollection = db.Collection("customers")
	AccountCollection = db.Collection("accounts") 
}