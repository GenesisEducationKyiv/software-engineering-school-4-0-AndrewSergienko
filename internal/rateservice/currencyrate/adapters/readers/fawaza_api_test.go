package readers

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/rateservice/currencyrate/infrastructure"
	"testing"
)

type FawazaAPIAdapterSuite struct {
	suite.Suite
	adapter *FawazaAPICurrencyReader
}

func (suite *FawazaAPIAdapterSuite) SetupSuite() {
	settings := infrastructure.GetCurrencyAPISettings()
	suite.adapter = NewFawazaAPICurrencyReader(settings.FawazaAPIURL)
	suite.NotNil(suite.adapter)
}

func (suite *FawazaAPIAdapterSuite) TestGetCurrencyRate() {
	result, err := suite.adapter.GetCurrencyRate("USD", "UAH")

	suite.Nil(err)
	suite.NotEqual(0, result)
}

func TestFawazaAPIAdapterSuite(t *testing.T) {
	suite.Run(t, new(FawazaAPIAdapterSuite))
}
