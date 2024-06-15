package adapters

import (
	"fmt"
	"go_service/internal/infrastructure"
	"log"
	"net/smtp"
)

type EmailAdapter struct {
	Username string
	Auth     smtp.Auth
}

func GetEmailAdapter(settings infrastructure.EmailSettings) EmailAdapter {
	return EmailAdapter{
		Username: settings.Email,
		Auth:     smtp.PlainAuth("", settings.Email, settings.Password, settings.Host),
	}
}

func (ea EmailAdapter) Send(target string, rate float32) {
	to := []string{target}

	msg := []byte(
		"Subject: USD rate\r\n" +
			"\r\n" +
			"USD rate: " + fmt.Sprintf("%f", rate) + "\r\n",
	)

	err := smtp.SendMail("smtp.gmail.com:587", ea.Auth, ea.Username, to, msg)
	if err != nil {
		log.Println("Send email error:", err)
	}
}
