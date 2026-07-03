package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type ValidateDTO struct {
	CardNumber      string
	HolderName      string
	ExpirationMonth string
	ExpirationYear  string
	CVV             string
}

func GetToken(validateDTO ValidateDTO, keySecret string) (string, error) {
	cardData := strings.Join([]string{
		validateDTO.CardNumber,
		validateDTO.HolderName,
		validateDTO.ExpirationMonth,
		validateDTO.ExpirationYear,
		validateDTO.CVV,
	}, "|")

	mac := hmac.New(sha256.New, []byte(keySecret))
	if _, err := mac.Write([]byte(cardData)); err != nil {
		return "", err
	}

	hashBytes := mac.Sum(nil)
	token := base64.StdEncoding.EncodeToString(hashBytes)

	return token, nil
}
