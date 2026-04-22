package model

import (
	"adquirer/entity"
	"adquirer/repository"
)

func CardValuesByCard(cardDescription string) (*entity.Card, error) {
	card, err := repository.CardValuesByCard(cardDescription)

	return card, err
}
