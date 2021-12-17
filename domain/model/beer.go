package model

type Beer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Brewery  string
	Country  string
	Price    int64
	Currency string
}
