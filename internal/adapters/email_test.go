package adapters

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go_service/internal/common"
	"go_service/internal/infrastructure"
	"log"
	"testing"
)

type EmailAdapterTestSuite struct {
	suite.Suite
	adapter   common.EmailSender
	container testcontainers.Container
	ctx       context.Context
}

func (suite *EmailAdapterTestSuite) SetupSuite() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mailhog/mailhog",
		ExposedPorts: []string{"1025/tcp", "8025/tcp"},
		WaitingFor:   wait.ForListeningPort("1025/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	suite.NoError(err)

	suite.ctx = ctx
	suite.container = container

	host, err := container.Host(ctx)
	suite.NoError(err)
	smtpPort, err := container.MappedPort(ctx, "1025")
	suite.NoError(err)

	smtpAddress := fmt.Sprintf("%s:%s", host, smtpPort.Port())
	suite.NotNil(smtpAddress)

	settings := infrastructure.GetEmailSettings()
	settings.Host = smtpAddress
	settings.Email = "sender@test.com"
	settings.Password = ""
	suite.adapter = GetEmailAdapter(settings)
}

func (suite *EmailAdapterTestSuite) TearDownSuite() {
	err := suite.container.Terminate(suite.ctx)
	if err != nil {
		log.Println(err)
	}
}

func (suite *EmailAdapterTestSuite) TestSend() {
	suite.NoError(suite.adapter.Send("test@gmail.com", 100))
}

func TestEmailAdapterTestSuite(t *testing.T) {
	suite.Run(t, new(EmailAdapterTestSuite))
}
