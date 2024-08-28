package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/repository/postgres"
)

type AuthRouter struct {
}

func NewAuthRouter(db *postgres.Database) API {
	return &AuthRouter{}
}

func (a *AuthRouter) Setup(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "login"})
	})

}
