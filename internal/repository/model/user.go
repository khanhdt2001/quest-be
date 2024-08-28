package model

type LoginType string

const (
	PASSWORD       LoginType = "PASSWORD"
	GOOGLE_OAUTH   LoginType = "GOOGLE_OAUTH"
	FACEBOOK_OAUTH LoginType = "FACEBOOK_OAUTH"
)

type User struct {
	Id             int64     `gorm:"primaryKey" json:"id"`
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	PassWordHashed string    `json:"-"`
	LastLoginType  LoginType `gorm:"column:last_login_type;default:PASSWORD" sql:"type:login_type" json:"last_login_type"`
}
