package authentication

import (
	"errors"
	"fmt"
	"tracking_backend/src/dtos"
	dto "tracking_backend/src/dtos"
	"tracking_backend/src/models"
	validate "tracking_backend/src/validator"
)

// AuthenticationService handles the logic for user authentication
type AuthenticationService struct {
	UserModel *models.User_Model
}

// Creates a new instance
func NewAuthenticationService(user_model *models.User_Model) *AuthenticationService {
	return &AuthenticationService{UserModel: user_model}
}

// Login
func (service *AuthenticationService) Login(req dto.LoginRequest) (*dtos.User, error) {
	err := validate.LoginRequest(req)
	if err != nil {
		return nil, err
	}

	user, err := service.UserModel.LoginUser(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Register
func (service *AuthenticationService) Register(req dto.RegistrationRequest) error {
	fmt.Println("Starting registration function for " + req.Email + "...")

	exists, err := service.UserModel.CheckUserExists(req.Email)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("\tRegistration failed - an account with this email already exists.")
		return errors.New("email is already in use")
	}

	fmt.Println("\tCreating new user...")
	err = service.UserModel.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return err
	}

	return nil
}
