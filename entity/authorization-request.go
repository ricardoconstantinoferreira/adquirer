package entity

type AuthorizationRequest struct {
	Token    string          `json:"token"`
	Amount   float64         `json:"amount"`
}