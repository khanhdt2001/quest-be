package util

import (
	"github.com/quest-be/constant"
	"gopkg.in/mail.v2"
)

func SendMail(email string, subject string, body string) error {

	m := mail.NewMessage()
	m.SetHeader("From", constant.SENDER_EMAIL)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := mail.NewDialer(
		"smtp.gmail.com",
		587,
		constant.SENDER_EMAIL,
		Default.MAIL_PASSWORD)
	return d.DialAndSend(m)
}
