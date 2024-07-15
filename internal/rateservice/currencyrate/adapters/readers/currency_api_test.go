package readers

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/rateservice/currencyrate/infrastructure"
	"testing"
)

type CurrencyAPIAdapterSuite struct {
	suite.Suite
	adapter *CurrencyAPICurrencyReader
}

func (suite *CurrencyAPIAdapterSuite) SetupSuite() {
	settings := infrastructure.GetCurrencyAPISettings()
	suite.adapter = NewCurrencyAPICurrencyReader(settings.CurrencyAPIURL)
	suite.NotNil(suite.adapter)
}

func (suite *CurrencyAPIAdapterSuite) TestGetCurrencyRate() {
	result, err := suite.adapter.GetCurrencyRate("USD", "UAH")

	suite.Nil(err)
	suite.NotEqual(0, result)
}

func TestCurrencyAPIAdapterSuite(t *testing.T) {
	suite.Run(t, new(CurrencyAPIAdapterSuite))
}
