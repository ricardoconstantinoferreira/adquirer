package model

import "adquirer/repository"

func CardValuesUpdate(card string, requestValue float64, total float64) error {
	return repository.CardValuesUpdate(card, requestValue, total)
}
