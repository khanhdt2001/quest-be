package dto

type ResendOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
}
