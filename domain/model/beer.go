package model

type Beer struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewety"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}
