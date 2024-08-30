package dto

type CreateUserRequest struct {
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password"`
	LastLoginType string `json:"last_login_type" validate:"required,oneof=PASSWORD GOOGLE_OAUTH FACEBOOK_OAUTH"`
}
