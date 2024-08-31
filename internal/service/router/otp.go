package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/service/controller"
)

type OtpRouter struct {
	controller *controller.OtpController
}

func NewOtpRouter() *OtpRouter {
	return &OtpRouter{}
}

func (o *OtpRouter) SetController(controller *controller.OtpController) {
	o.controller = controller
}

func (o *OtpRouter) Setup(router *gin.RouterGroup) {
	router.POST("/resend-otp", o.controller.ResendOtp)
}
