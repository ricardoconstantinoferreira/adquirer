package entity

type CaptureRequest struct {
	Token 	string 	`json:"token,omitempty"`
	Amount	float64 `json:"amount,omitempty"`
}