package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/service/controller"
)

type AuthRouter struct {
	controller *controller.AuthController
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) SetController(controller *controller.AuthController) {
	a.controller = controller
}

func (a *AuthRouter) Setup(router *gin.RouterGroup) {

	router.POST("/signup", a.controller.SignUp)
	router.POST("/verify", a.controller.VerifyUser)
}
