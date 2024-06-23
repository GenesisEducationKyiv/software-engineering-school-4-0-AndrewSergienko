package currency_rate

type APIReaderFacade struct {
	readers []APICurrencyReader
}

func NewAPIReaderFacade(readers []APICurrencyReader) *APIReaderFacade {
	return &APIReaderFacade{readers: readers}
}

func (facade *APIReaderFacade) GetCurrencyRate(from string, to string) (float32, error) {
	var rate float32
	var err error

	for _, reader := range facade.readers {
		rate, err = reader.GetCurrencyRate(from, to)
		if err == nil {
			return rate, nil
		}
	}
	return 0, err
}
