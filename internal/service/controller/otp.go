package controller

import "github.com/quest-be/internal/service/handler"

type OtpController struct {
	handler handler.IOtpHandler
}

func NewOtpController() *OtpController {
	return &OtpController{}
}

func (o *OtpController) SetHandler(handler handler.IOtpHandler) {
	o.handler = handler
}

func (o *OtpController) CreateOtp() {

}
