package main

import (
	"log"

	"gocashflow/api"
	"gocashflow/config"
)

func main() {
	config.ConnectToMongo()

	app := api.NewServer()

	err := app.ListenTLS(":3000", "./cert.pem", "./cert.key")
	if err != nil {
		log.Fatal("❌ Failed to start TLS server:", err)
	}
}