package repository

import (
	"adquirer/db"
	"adquirer/entity"
)

func CardValuesByCard(description string) (*entity.Card, error) {

	conn := db.ReturnDb()

	query := "select id, card, flag, cvv, venc, total from card_adquirer where card = ?"
	result := conn.QueryRow(query, description)

	card := &entity.Card{}

	err := result.Scan(&card.Id, &card.Card, &card.Flag, &card.Cvv, &card.Venc, &card.Total)

	if err != nil {
		return nil, err
	}

	return card, nil
}

func CardValuesUpdate(card string, requestTotal float64, total float64) error {
	result := total - requestTotal
	conn := db.ReturnDb()

	query := "update card_adquirer set total = ? where card = ?"
	_, err := conn.Exec(query, result, card)

	if err != nil {
		return err
	}

	return nil
}
