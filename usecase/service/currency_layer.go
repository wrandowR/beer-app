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

func NewCurrencyLayer() {
	//ExternalBilling.(*ExternalBillingService).Client = resty.New().SetRetryCount(config.RetryCount)
}

//Convert one currency to another
func (c *CurrencyLayer) Convert(fromCurrency string, toCurrency string, amount string) error {
	url := config.CurrencyLayerURL() + "live?access_key=" + config.CurrencyLayerAPIKEY() + "&currencies=USD,AUD,CAD,PLN,MXN&format=1"

	response, err := c.Client.R().Get(url)
	if err != nil {
		return merry.Wrap(err)
	}

	code := response.StatusCode()
	if code != http.StatusOK && code != http.StatusCreated {
		return merry.New("convert one currency to another").WithValue("http status: ", response.StatusCode())
	}

	var currencyResponse contract.CurrencyLayerResponse

	err = json.Unmarshal(response.Body(), &currencyResponse)
	if err != nil {
		return merry.Wrap(err)
	}

	fmt.Println("response", currencyResponse)

	return nil

}
