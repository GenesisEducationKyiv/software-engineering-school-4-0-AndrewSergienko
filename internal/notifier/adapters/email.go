package adapters

import (
	"fmt"
	"go_service/internal/infrastructure"
	"net/smtp"
)

type EmailAdapter struct {
	username string
	host     string
	auth     smtp.Auth
}

func NewEmailAdapter(settings infrastructure.EmailSettings) EmailAdapter {
	return EmailAdapter{
		username: settings.Email,
		host:     settings.Host,
		auth:     smtp.PlainAuth("", settings.Email, settings.Password, settings.Host),
	}
}

func (ea EmailAdapter) Send(target string, rate float32) error {

	to := []string{target}
	subject := "Subject: USD rate\r\n"
	from := "From: " + ea.username + "\r\n"
	toHeader := "To: target@example.com\r\n"
	body := "USD rate: " + fmt.Sprintf("%f", rate) + "\r\n"

	msg := []byte(from + toHeader + subject + "\r\n" + body)

	return smtp.SendMail(ea.host, nil, ea.username, to, msg)
}
