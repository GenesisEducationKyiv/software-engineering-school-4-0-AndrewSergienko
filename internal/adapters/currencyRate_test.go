package adapters

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/infrastructure"
	"testing"
)

type CurrencyRateAdapterSuite struct {
	suite.Suite
	adapter *APICurrencyReader
}

func (suite *CurrencyRateAdapterSuite) SetupSuite() {
	settings := infrastructure.GetCurrencyAPISettings()
	suite.adapter = NewAPICurrencyReader(settings)
}

func (suite *CurrencyRateAdapterSuite) TestGetUSDCurrencyRate() {
	result, err := suite.adapter.GetUSDCurrencyRate()

	suite.Nil(err)
	suite.NotEqual(0, result)
}

func TestCurrencyRateAdapterSuite(t *testing.T) {
	suite.Run(t, new(CurrencyRateAdapterSuite))
}
