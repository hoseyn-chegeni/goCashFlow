package main

import (
	"log"

	"gocashflow/api"
	"gocashflow/config" 
)

func main() {
	config.ConnectToMongo()
	app := api.NewServer()
	log.Fatal(app.Listen(":3000"))
}
