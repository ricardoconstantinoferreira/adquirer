package handler

import (
	"adquirer/model"
	"adquirer/validation"
	"encoding/json"
	"net/http"
)

type ValidationRequest struct {
	Card  string  `json:"card"`
	CVV   string  `json:"cvv"`
	Venc  string  `json:"venc"`
	Total float64 `json:"total"`
}

type ValidationResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func ValidateCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req ValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ValidationResponse{
			Message: "Payload inválido",
			Code:    "96",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if !validation.IsValidLuhn(req.Card) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ValidationResponse{
			Message: "Cartão inválido",
			Code:    "14",
		})
		return
	}

	resultValidCard, total := validation.ValidationTotal(req.Card, req.Total)

	if !resultValidCard && total == 0 {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ValidationResponse{
			Message: "Saldo insuficiente",
			Code:    "51",
		})
		return
	}

	error := model.CardValuesUpdate(req.Card, req.Total, total)

	if error != nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ValidationResponse{
			Message: "Erro ao alterar o saldo",
			Code:    "13",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ValidationResponse{
		Message: "Transacao autorizada com sucesso",
		Code:    "00",
	})
}
