package dto

type LoginByPasswordRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginByGoogleRequest struct {
	Token string `json:"token" binding:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=10"`
}

type VerifyUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,min=6,max=6,number"`
}

type GoogleOAuthBody struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RedirectUri  string `json:"redirect_uri"`
}

type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`
}
