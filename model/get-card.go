package model

import (
	"adquirer/entity"
	"adquirer/repository"
	"strings"
)

func CardValuesByCard(cardDescription string) (*entity.Card, error) {
	card, err := repository.CardValuesByCard(cardDescription)

	return card, err
}

func CardValuesByToken(token string) (*entity.Card, error) {
	card, err := repository.CardValueByToken(token)

	return card, err
}

func GetLastFourDigits(card string) string {
	clean := strings.NewReplacer(" ", "", "-", "").Replace(strings.TrimSpace(card))
	if len(clean) <= 4 {
		return clean
	}

	return clean[len(clean)-4:]
}

func ExtractExpiration(req entity.ValidationRequest) (string, string) {
	month := strings.TrimSpace(req.ExpirationMonth)
	year := strings.TrimSpace(req.ExpirationYear)

	if month != "" || year != "" {
		return month, year
	}

	venc := strings.TrimSpace(req.Venc)
	if venc == "" {
		return "", ""
	}

	parts := strings.Split(venc, "/")
	if len(parts) != 2 {
		return venc, ""
	}

	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}
