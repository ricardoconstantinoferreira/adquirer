package entity

type CaptureResponse struct {
	TransactionId 		string 		`json:"transactionId,omitempty"`
	Status 				string 		`json:"status,omitempty"`
	CapturedAmount 		float64 	`json:"capturedAmount,omitempty"`
	CapturedAt			string 		`json:"capturedAt,omitempty"`
	Message				string 		`json:"message,omitempty"`
}