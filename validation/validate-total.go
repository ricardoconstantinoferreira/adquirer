package validation

import "adquirer/model"

func ValidationTotal(card string, value float64) (bool, float64) {

	result, err := model.CardValuesByCard(card)

	if err != nil {
		return false, 0
	}

	if result.Total <= value {
		return false, 0
	}

	return true, result.Total
}
