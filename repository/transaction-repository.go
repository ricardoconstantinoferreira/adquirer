package repository

import (
	"adquirer/db"
	"adquirer/entity"
	"time"
)

func SaveTransaction(transaction_id string, card_adquirer_id int, amount float64) error {
	conn := db.ReturnDb()

	query := "insert into transaction (id, transaction_id, card_adquirer_id, captured_at, amount) values (?, ?, ?, ?, ?)"
	_, err := conn.Exec(query, 0, transaction_id, card_adquirer_id, time.Now(), amount)

	if err != nil {
		return err
	}

	return nil
}


func GetTransactionByTransactionId(transaction_id string) (*entity.Transaction, error) {

	conn := db.ReturnDb()

	query := "select id, transaction_id, card_adquirer_id, captured_at, amount from transaction where transaction_id = ?"
	result := conn.QueryRow(query, transaction_id)

	transaction := &entity.Transaction{}

	err := result.Scan(&transaction.Id, &transaction.TransactionId, &transaction.CardAdquirerId, &transaction.CapturedAt, &transaction.Amount)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}