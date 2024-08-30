package constant

import "errors"

var (
	ErrUserAlreadyExist    = errors.New("user already exist")
	ErrInvalidToken        = errors.New("token is invalid")
	ErrExpiredToken        = errors.New("token has expired")
	ErrUserAlreadyVerified = errors.New("user already verified")
	ErrInvalidOTP          = errors.New("invalid otp")
	ErrExpiredOTP          = errors.New("otp has expired")
)
