package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/service/dto"
	"github.com/quest-be/internal/service/handler"
)

type OtpController struct {
	handler handler.IOtpHandler
}

func NewOtpController() *OtpController {
	return &OtpController{}
}

func (o *OtpController) SetHandler(handler handler.IOtpHandler) {
	o.handler = handler
}

func (o *OtpController) ResendOtp(c *gin.Context) {
	var req dto.ResendOtpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := o.handler.ResendOtp(c, req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
