package dtos

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,validateEmail"`
	Password string `json:"password" validate:"required,min=6"`
}
