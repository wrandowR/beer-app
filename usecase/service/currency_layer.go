package service

import (
	"ZachIgarz/test-beer/config"
	"ZachIgarz/test-beer/domain/contract"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ansel1/merry"
	"github.com/go-resty/resty/v2"
)

type CurrencyLayer struct {
	Client *resty.Client
}

var CurrencyLayerService contract.CurrencyLayer = &CurrencyLayer{}

func NewCurrencyLayer() {
	CurrencyLayerService.(*CurrencyLayer).Client = resty.New().SetRetryCount(5)
}

//Convert one currency to another
func (c *CurrencyLayer) Convert(fromCurrency string, toCurrency string) (float32, error) {
	url := config.CurrencyLayerURL() + "live?access_key=" + config.CurrencyLayerAPIKEY() + "&currencies=USD,AUD,CAD,PLN,MXN&format=1"

	fmt.Println(url)
	response, err := c.Client.R().Get(url)
	if err != nil {
		return 0, merry.Wrap(err)
	}

	code := response.StatusCode()
	if code != http.StatusOK && code != http.StatusCreated {
		return 0, merry.New("convert one currency to another").WithValue("http status: ", response.StatusCode())
	}

	var currencyResponse contract.CurrencyLayerResponse

	err = json.Unmarshal(response.Body(), &currencyResponse)
	if err != nil {
		return 0, merry.Wrap(err)
	}

	for key, val := range currencyResponse.Quotes {
		if key == fromCurrency+toCurrency {
			return val, nil
		}
	}

	return 0, nil
}
