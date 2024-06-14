package interfaces

import "github.com/gorilla/mux"

// Requires a 'InitializeRoutes' method for configuring routes
type Controller interface {
	InitializeRoutes(router *mux.Router) // Sets up the routes for this controller
	GetBasePath() string                 // Returns the base path of the controller
}
