package model

import "adquirer/repository"

func CardValuesUpdate(card string, requestValue float64, total float64) error {
	return repository.CardValuesUpdate(card, requestValue, total)
}

func CardValueTokenUpdate(token string, card string) error {
	return repository.CardValueToken(token, card)
}

func CardValueCaptureUpdate(token string) error {
	return repository.CardValueCapture(token)
}
