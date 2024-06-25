package validate

import (
	"regexp"
	"tracking_backend/src/dtos"

	"github.com/go-playground/validator/v10"
)

// Validate is a package-level variable that holds the shared instance of the validator
var validate *validator.Validate

func init() {
	// Initialize the validator instance when the package is imported
	validate = validator.New()
	validate.RegisterValidation("validateEmail", Email)
}

func Email(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,8}$`).MatchString(fl.Field().String())
}

// Validates the LoginRequest fields
func LoginRequest(req dtos.LoginRequest) error {
	// Use the validator to check the struct based on the tags
	return validate.Struct(req)
}

// Validates the RegistrationRequest fields
func RegistrationRequest(req dtos.RegistrationRequest) error {
	// Use the validator to check the struct based on the tags
	return validate.Struct(req)
}
