package authentication

import (
	"encoding/json"
	"fmt"
	dto "linkpulse_api/src/api/dtos"
	"linkpulse_api/src/database"
	"linkpulse_api/src/interfaces"
	"net/http"

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

	// Create a subrouter for this controller based on its base path
	subRouter := router.PathPrefix(ac.GetBasePath()).Subrouter()

	subRouter.HandleFunc("/register", ac.registerHandler).Methods("POST")

	/*
		// Define routes and their corresponding handler functions
		subRouter.HandleFunc("/login", ac.Service.Login).Methods("POST")
		subRouter.HandleFunc("/register", ac.Service.Register).Methods("POST")
		subRouter.HandleFunc("/reset-password", ac.Service.ResetPassword).Methods("POST")
		subRouter.HandleFunc("/validate", ac.Service.Validate).Methods("GET")


	*/
}

func (ac *AuthenticationController) registerHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := dto.ValidateRegistrationRequest(req); err != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
		return
	}

	err := ac.Service.Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User successfully created")
}
