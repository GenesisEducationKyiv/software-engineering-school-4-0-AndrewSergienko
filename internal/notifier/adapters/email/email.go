package email

import (
	"fmt"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/metrics"
	"net/smtp"
)

type Adapter struct {
	username string
	host     string
	auth     smtp.Auth
}

func NewEmailAdapter(settings infrastructure.EmailSettings) Adapter {
	return Adapter{
		username: settings.Email,
		host:     settings.Host,
		auth:     smtp.PlainAuth("", settings.Email, settings.Password, settings.Host),
	}
}

func (ea Adapter) Send(target string, rate float32) error {

	to := []string{target}
	subject := "Subject: USD rate\r\n"
	from := "From: " + ea.username + "\r\n"
	toHeader := "To: " + target + "\r\n"
	body := "USD rate: " + fmt.Sprintf("%f", rate) + "\r\n"

	msg := []byte(from + toHeader + subject + "\r\n" + body)

	err := smtp.SendMail(ea.host, nil, ea.username, to, msg)

	var status string
	if err != nil {
		status = "error"
	} else {
		status = "success"
	}
	metrics.EmailsSentTotal.WithLabelValues(status).Inc()
	return err
}
