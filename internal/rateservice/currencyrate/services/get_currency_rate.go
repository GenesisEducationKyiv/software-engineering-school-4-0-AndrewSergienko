package services

type GetCurrencyRateInputDTO struct {
	From string
	To   string
}

type GetCurrencyRateOutputDTO struct {
	Result float32
	Err    error
}

type CurrencyGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type GetCurrencyRate struct {
	currencyGateway CurrencyGateway
}

func NewGetCurrencyRate(currencyGateway CurrencyGateway) *GetCurrencyRate {
	return &GetCurrencyRate{currencyGateway: currencyGateway}
}

func (s *GetCurrencyRate) Handle(data GetCurrencyRateInputDTO) GetCurrencyRateOutputDTO {
	result, err := s.currencyGateway.GetCurrencyRate(data.From, data.To)
	return GetCurrencyRateOutputDTO{Result: result, Err: err}
}
