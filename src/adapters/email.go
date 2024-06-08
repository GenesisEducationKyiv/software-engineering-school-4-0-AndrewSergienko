package adapters

import (
	"fmt"
	"net/smtp"
)

type EmailAdapter struct {
	Username string
	Password string
	Auth     smtp.Auth
}

func (ea EmailAdapter) CreateAuth() {
	ea.Auth = smtp.PlainAuth("", ea.Username, ea.Password, "smtp.gmail.com")
}

func (ea EmailAdapter) Send(target string, rate float32) {
	println(fmt.Sprintf("Send email from: %s; to: %s; rate: %f", ea.Username, target, rate))
	/*
		to := []string{target}

		msg := []byte("To: kate.doe@example.com\r\n" +
			"Subject: USD rate\r\n" +
			"\r\n" +
			"USD rate: " + fmt.Sprintf("%f", rate) + "\r\n")

		err := smtp.SendMail("smtp.gmail.com:587", ea.Auth, ea.Username, to, msg)

		if err != nil {
			log.Fatal(err)
		}
	*/
}
