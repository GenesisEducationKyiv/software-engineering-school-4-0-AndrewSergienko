package services

type GetCurrencyRateInputDTO struct {
	from string
	to   string
}

type CurrencyGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type GetCurrencyRate struct {
	currencyGateway CurrencyGateway
}

func (s *GetCurrencyRate) Handle(data GetCurrencyRateInputDTO) (float32, error) {
	return s.currencyGateway.GetCurrencyRate(data.from, data.to)
}
