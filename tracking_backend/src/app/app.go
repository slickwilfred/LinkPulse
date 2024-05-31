package app

import (
	"tracking_backend/src/database"
	"tracking_backend/src/interfaces"

	"github.com/gorilla/mux"
)

// Holds the router and database instance
type App struct {
	Router *mux.Router
	DB     *database.DB
}

// Initializes a new app
func NewApp(db *database.DB) *App {
	return &App{
		Router: mux.NewRouter(),
		DB:     db,
	}
}

// Registers a controller and its routes
func (a *App) RegisterControllers(controllers []interfaces.Controller) {
	for _, controller := range controllers {
		controller.InitializeRoutes(a.Router)
	}
}

// Returns the App's router
func (a *App) GetRouter() *mux.Router {
	return a.Router
}
