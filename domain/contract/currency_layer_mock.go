package contract

import "github.com/stretchr/testify/mock"

type CurrencyLayerMock struct {
	mock.Mock
}

func (c *CurrencyLayerMock) Convert(fromCurrency string, toCurrency string) (float32, error) {
	args := c.Called(fromCurrency, toCurrency)

	return args.Get(0).(float32), args.Error(1)
}
