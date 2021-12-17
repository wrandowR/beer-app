package contract

type CurrencyLayer interface {
	Convert(fromCurrency string, toCurrency string) (float32, error)
}

type CurrencyLayerResponse struct {
	Success   bool               `json:"success"`
	Source    string             `json:"source"`
	Timestamp int64              `json:"timestamp"`
	Quotes    map[string]float32 `json:"quotes"`
}
