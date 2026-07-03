package entity

type ValidationRequest struct {
	Card            string  `json:"card"`
	CVV             string  `json:"cvv"`
	Venc            string  `json:"venc"`
	HolderName      string  `json:"holderName"`
	ExpirationMonth string  `json:"expirationMonth"`
	ExpirationYear  string  `json:"expirationYear"`
	Total           float64 `json:"total"`
}