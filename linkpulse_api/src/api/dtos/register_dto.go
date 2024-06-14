package dto

type RegistrationRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,validateEmail"`
	Password string `json:"password" validate:"required,min=6"`
}
