package email

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
)

type EmailAdapterTestSuite struct {
	suite.Suite
	adapter Adapter
}

func (suite *EmailAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetEmailSettings()
	settings.Host = "localhost:1025"
	settings.Email = "sender@test.com"
	settings.Password = ""
	suite.adapter = NewEmailAdapter(settings)
}

func (suite *EmailAdapterTestSuite) TestSend() {
	err := suite.adapter.Send("test@gmail.com", 100)
	var netOpError *net.OpError
	if err != nil && errors.As(err, &netOpError) {
		suite.T().Skip()
	}
	suite.NoError(err)
	suite.True(checkMail("sender@test.com", "test@gmail.com"))
}

func TestEmailAdapterTestSuite(t *testing.T) {
	suite.Run(t, new(EmailAdapterTestSuite))
}

type Mailbox struct {
	Mailbox string `json:"Mailbox"`
	Domain  string `json:"Domain"`
}

type EmailAddress struct {
	Mailbox string `json:"Mailbox"`
	Domain  string `json:"Domain"`
	Params  string `json:"Params"`
}

type ContentHeaders struct {
	From       []string `json:"From"`
	MessageID  []string `json:"Message-ID"`
	Received   []string `json:"Received"`
	ReturnPath []string `json:"Return-Path"`
	Subject    []string `json:"Subject"`
	To         []string `json:"To"`
}

type EmailContent struct {
	Headers ContentHeaders `json:"Headers"`
	Body    string         `json:"Body"`
	Size    int            `json:"Size"`
}

type EmailItem struct {
	ID      string         `json:"ID"`
	From    EmailAddress   `json:"From"`
	To      []EmailAddress `json:"To"`
	Content EmailContent   `json:"Content"`
}

type MailhogResponse struct {
	Total int         `json:"total"`
	Count int         `json:"count"`
	Start int         `json:"start"`
	Items []EmailItem `json:"items"`
}

func checkMail(from string, to string) (bool, error) {
	resp, err := http.Get("http://127.0.0.1:8025/api/v2/messages")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var mailhogResponse MailhogResponse
	if err = json.Unmarshal(body, &mailhogResponse); err != nil {
		return false, err
	}

	for _, item := range mailhogResponse.Items {
		fromHeader := strings.Join(item.Content.Headers.From, ",")
		if strings.Contains(fromHeader, from) {
			toHeader := strings.Join(item.Content.Headers.To, ",")
			if strings.Contains(toHeader, to) {
				return true, nil
			}
		}
	}
	return false, nil
}
