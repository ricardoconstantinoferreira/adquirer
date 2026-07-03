package entity

type ValidationResponse struct {
	Message         string `json:"message"`
	Code            string `json:"code"`
	Brand           string `json:"brand"`
	CardToken       string `json:"cardToken"`
	ExpirationMonth string `json:"expirationMonth"`
	ExpirationYear  string `json:"expirationYear"`
	LastFourDigits  string `json:"lastFourDigits"`
}