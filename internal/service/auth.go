package service

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/database"
)

type AuthController struct {
}

func NewAuthController(db *database.Database) API {
	return &AuthController{}
}

func (a *AuthController) Setup(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "login"})
	})

}
