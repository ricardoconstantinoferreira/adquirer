package model

import (
	"adquirer/entity"
	"adquirer/repository"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func Save(token string, amount float64) (*entity.CaptureResponse, error) {
	result, err := CardValuesByToken(token)

	if err != nil {
		return nil, err
	}

	res := createTransactionHash(token)
	error := repository.SaveTransaction(res, result.Id, amount)

	if error != nil {
		return nil, error
	}

	transaction, err := GetTransactionByTransaction(res)

	if err != nil {
		return nil, err
	}

	return &entity.CaptureResponse{
		TransactionId: transaction.TransactionId,
		Status: "paid",
		CaptureAmount: amount,
		CapturedAt: transaction.CapturedAt.Format(time.RFC3339),
		Message: "Captura realizada com sucesso.",
	}, nil
}

func createTransactionHash(token string) string {
	payload := fmt.Sprintf("%s|%s", token, time.Now().Format(time.RFC3339Nano))
	sum := sha256.Sum256([]byte(payload))
	return base64.StdEncoding.EncodeToString(sum[:])
}