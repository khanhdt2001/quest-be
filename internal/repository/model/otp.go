package model

import "time"

type Otp struct {
	Id        uint64    `gorm:"primaryKey" json:"id"`
	UserId    uint64    `json:"user_id"`
	OTP       string    `json:"otp"`
	ExpiredAt time.Time `json:"expired_at"`
	User      User
}
