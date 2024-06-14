package authentication

import (
	"tracking_backend/src/database"
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

// ResetPassword

// Validate
