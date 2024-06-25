package authentication

import (
	"net/http"
	"tracking_backend/src/dtos"
	"tracking_backend/src/interfaces"
	"tracking_backend/src/models"
	validate "tracking_backend/src/validator"

	"github.com/gin-gonic/gin"
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

// InitializeRoutes configures authentication routes
func (ac *AuthenticationController) InitializeRoutes(router *gin.Engine) {
	authGroup := router.Group(ac.GetBasePath())
	{
		authGroup.POST("/register", ac.registerHandler)
		authGroup.POST("/login", ac.loginHandler)
	}
}

func (ac *AuthenticationController) loginHandler(c *gin.Context) {
	var req dtos.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request."})
		return
	}

	user, err := ac.Service.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dtos.LoginResponse{
		User:    user,
		Message: "Login successful",
	}

	c.JSON(http.StatusAccepted, response)
}

// Registration Handler
func (ac *AuthenticationController) registerHandler(c *gin.Context) {
	var req dtos.RegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validate.RegistrationRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.Service.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully created"})
}
