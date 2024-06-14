package authentication

import (
	"tracking_backend/src/database"
	"tracking_backend/src/interfaces"

	"github.com/gorilla/mux"
)

type AuthenticationController struct {
	Service  *AuthenticationService
	basePath string
}

// Creates a new instance of AuthenticationController
func NewAuthenticationController(db *database.DB) interfaces.Controller {
	return &AuthenticationController{
		Service:  NewAuthenticationService(db),
		basePath: "/auth",
	}
}

// GetBasePath returns the base path for this controller
func (ac *AuthenticationController) GetBasePath() string {
	return ac.basePath
}

// IntializeRoutes configures authentication routes
func (ac *AuthenticationController) InitializeRoutes(router *mux.Router) {

	/*
		// Create a subrouter for this controller based on its base path
		//subRouter := router.PathPrefix(ac.GetBasePath()).Subrouter()

		// Define routes and their corresponding handler functions
		subRouter.HandleFunc("/login", ac.Service.Login).Methods("POST")
		subRouter.HandleFunc("/register", ac.Service.Register).Methods("POST")
		subRouter.HandleFunc("/reset-password", ac.Service.ResetPassword).Methods("POST")
		subRouter.HandleFunc("/validate", ac.Service.Validate).Methods("GET")


	*/
}
