package authentication

import (
	"encoding/json"
	"fmt"
	"linkpulse_api/src/dtos"
	"linkpulse_api/src/interfaces"
	"linkpulse_api/src/models"
	validate "linkpulse_api/src/validator"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthenticationController struct {
	Service  *AuthenticationService
	basePath string
}

// Creates a new instance of AuthenticationController
func NewAuthenticationController(user_model *models.User_Model) interfaces.Controller {
	return &AuthenticationController{
		Service:  NewAuthenticationService(user_model),
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
	subRouter.HandleFunc("/login", ac.loginHandler).Methods("POST")
	/*
		// Define routes and their corresponding handler functions
		subRouter.HandleFunc("/login", ac.Service.Login).Methods("POST")
		subRouter.HandleFunc("/register", ac.Service.Register).Methods("POST")
		subRouter.HandleFunc("/reset-password", ac.Service.ResetPassword).Methods("POST")
		subRouter.HandleFunc("/validate", ac.Service.Validate).Methods("GET")


	*/
}

func (ac *AuthenticationController) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req dtos.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request."})
		return
	}

	user, err := ac.Service.Login(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(fmt.Errorf("Login failed: %s", err))
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Login failed: %s", err)})
		return
	}

	response := dtos.LoginResponse{
		User:    user,
		Message: "Login successful",
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

// Registration Handler
func (ac *AuthenticationController) registerHandler(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.RegistrationRequest(req); err != nil {
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
