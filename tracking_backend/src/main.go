package main

import (
	"fmt"
	"log"
	"tracking_backend/src/app"
	"tracking_backend/src/controllers/authentication"
	"tracking_backend/src/database"
	"tracking_backend/src/interfaces"
	"tracking_backend/src/models"
)

func main() {
	// Load the database config file
	config, err := database.LoadConfig("./database/dbconfig.json")
	if err != nil {
		log.Fatalf("failed to load the db config file: %v", err)
	}

	// Init config file
	dbPool, err := database.Initialize(config)
	if err != nil {
		log.Fatalf("failed to initialized the db connection: %f", err)
	}

	// Init db
	db := database.NewDB(dbPool)

	// Init app (router/db - see struct)
	app := app.NewApp(db)

	// Init models

	// Add list of controllers
	controllers := []interfaces.Controller{
		authentication.NewAuthenticationController(&models.User_Model{}),
	}

	// Register controllers
	app.RegisterControllers(controllers)

	// Start the server
	port := ":8888"
	listen := fmt.Sprintf("\nServer listening on port %s...", port)
	log.Println(listen)
	log.Fatal(app.GetRouter().Run(port))
}
