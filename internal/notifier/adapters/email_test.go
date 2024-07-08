package adapters

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure"
	"net"
	"testing"
)

type EmailAdapterTestSuite struct {
	suite.Suite
	adapter EmailAdapter
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
}

func TestEmailAdapterTestSuite(t *testing.T) {
	suite.Run(t, new(EmailAdapterTestSuite))
}
