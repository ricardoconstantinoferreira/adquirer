package model

import (
	"adquirer/entity"
	"adquirer/repository"
)

func GetTransactionByTransaction(transaction_id string) (*entity.Transaction, error) {
	return repository.GetTransactionByTransactionId(transaction_id)
}