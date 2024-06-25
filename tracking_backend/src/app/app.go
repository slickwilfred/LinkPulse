package app

import (
	"tracking_backend/src/database"
	"tracking_backend/src/interfaces"

	"github.com/gin-gonic/gin"
)

// Holds the router and database instance
type App struct {
	Router *gin.Engine
	DB     *database.DB
}

// Initializes a new app
func NewApp(db *database.DB) *App {
	return &App{
		Router: gin.Default(),
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
func (a *App) GetRouter() *gin.Engine {
	return a.Router
}
