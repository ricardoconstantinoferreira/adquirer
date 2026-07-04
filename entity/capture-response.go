package entity

type CaptureResponse struct {
	TransactionId 		string 		`json:"transaction_id,omitempty"`
	Status 				string 		`json:"status,omitempty"`
	CaptureAmount 		float64 	`json:"capture_amount,omitempty"`
	CapturedAt			string 		`json:"captured_at,omitempty"`
	Message				string 		`json:"message,omitempty"`
}