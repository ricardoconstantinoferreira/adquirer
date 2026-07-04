package entity

import "time"

type Transaction struct {
	Id				int 		`json:"id"`
	TransactionId   string 		`json:"transaction_id"`
	CardAdquirerId	int 		`json:"card_adquirer_id"`
	CapturedAt		time.Time 		`json:"captured_at"`
	Amount			float64 		`json:"amount"`
}