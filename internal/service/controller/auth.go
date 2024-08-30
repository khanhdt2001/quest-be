package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/service/dto"
	"github.com/quest-be/internal/service/handler"
)

type AuthController struct {
	handler handler.IAuthHandler
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) SetHandler(handler handler.IAuthHandler) {
	a.handler = handler
}

func (a *AuthController) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.handler.SignUp(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (a *AuthController) VerifyUser(c *gin.Context) {
	// get otp from request
	// get email from request
	// get user by email
	// check if otp is correct
	// if correct, update user status to active
	// if not, return error
	var req dto.VerifyUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.handler.VerifyUser(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
