package contract

type CurrencyLayer interface {
	Convert(fromCurrency string, toCurrency string, amount string) (float64, error)
}

type CurrencyLayerResponse struct {
	Success   bool   `json:"success"`
	Source    string `json:"source"`
	Timestamp int64  `json:"timestamp"`
	Quotes    Quotes `json:"quotes"`
}

type Quotes struct {
	USDUSD float64 `json:"USDUSD"`
	USDAUD float64 `json:"USDAUD"`
	USDCAD float64 `json:"USDCAD"`
	USDPLN float64 `json:"USDPLN"`
	USDMXN float64 `json:"USDMXN"`
}
