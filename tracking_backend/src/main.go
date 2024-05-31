package main

import (
	"fmt"
	"log"
	"net/http"
	"tracking_backend/src/api/controllers/authentication"
	"tracking_backend/src/app"
	"tracking_backend/src/database"
	"tracking_backend/src/interfaces"
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

	// Add list of controllers
	controllers := []interfaces.Controller{
		authentication.NewAuthenticationController(db),
	}

	// Register controllers
	app.RegisterControllers(controllers)

	// Start the server
	port := ":8888"
	listen := fmt.Sprintf("\nServer listening on port %s...", port)
	log.Println(listen)
	log.Fatal(http.ListenAndServe(port, app.GetRouter()))
}
