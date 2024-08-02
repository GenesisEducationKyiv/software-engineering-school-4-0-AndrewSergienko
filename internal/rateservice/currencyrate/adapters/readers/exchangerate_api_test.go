package readers

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/rateservice/infrastructure"
	"testing"
)

type ExchangerateAPIAdapterSuite struct {
	suite.Suite
	adapter *ExchangerateAPICurrencyReader
}

func (suite *ExchangerateAPIAdapterSuite) SetupSuite() {
	settings := infrastructure.GetCurrencyAPISettings()
	suite.adapter = NewExchangerateAPICurrencyReader(settings.ExchangerateAPIURL)
	suite.NotNil(suite.adapter)
}

func (suite *ExchangerateAPIAdapterSuite) TestGetCurrencyRate() {
	result, err := suite.adapter.GetCurrencyRate("USD", "UAH")

	suite.Nil(err)
	suite.NotEqual(0, result)
}

func TestExchangerateAPIAdapterSuite(t *testing.T) {
	suite.Run(t, new(ExchangerateAPIAdapterSuite))
}
