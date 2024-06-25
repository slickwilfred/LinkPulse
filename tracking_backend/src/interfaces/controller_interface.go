package interfaces

import "github.com/gin-gonic/gin"

// Requires a 'InitializeRoutes' method for configuring routes
type Controller interface {
	InitializeRoutes(router *gin.Engine) // Sets up the routes for this controller
	GetBasePath() string                 // Returns the base path of the controller
}
