package authentication

import (
	"errors"
	"fmt"
	dto "linkpulse_api/src/api/dtos"
	"linkpulse_api/src/database"
	//"AffiliateLinksBackend/models/dto"
	//"context"
)

// AuthenticationService handles the logic for user authentication
type AuthenticationService struct {
	DB *database.DB
}

// Creates a new instance
func NewAuthenticationService(db *database.DB) *AuthenticationService {
	return &AuthenticationService{DB: db}
}

// Authenticate

// CreateCookie

// CreateToken

// Login

// Logout

// Register
func (service *AuthenticationService) Register(req dto.RegistrationRequest) error {
	fmt.Println("Starting registration function for " + req.Email + "...")

	exists, err := service.DB.CheckUserExists(req.Email)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("\tRegistration failed - an account with this email already exists.")
		return errors.New("email is already in use")
	}

	fmt.Println("\tCreating new user...")
	err = service.DB.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return err
	}

	return nil
}

// ResetPassword

// ValidateSession
