package handler

import (
	"adquirer/entity"
	"adquirer/model"
	"adquirer/validation"
	"encoding/json"
	"net/http"
	"os"
)

func ValidateCardTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req entity.ValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(entity.ValidationResponse{
			Message:         "Payload inválido",
			Code:            "96",
			Brand:           "",
			CardToken:       "",
			ExpirationMonth: "",
			ExpirationYear:  "",
			LastFourDigits:  "",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	expirationMonth, expirationYear := model.ExtractExpiration(req)
	responseBase := entity.ValidationResponse{
		Brand:           "",
		CardToken:       "",
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		LastFourDigits:  model.GetLastFourDigits(req.Card),
	}

	cardToken, err := model.GetToken(model.ValidateDTO{
		CardNumber:      req.Card,
		HolderName:      req.HolderName,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		CVV:             req.CVV,
	}, os.Getenv("CARD_TOKEN_SECRET"))
	if err != nil {
		w.WriteHeader(http.StatusOK)
		response := responseBase
		response.Message = "Erro ao gerar token"
		response.Code = "96"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	brand := model.CardFlagByNumber(req.Card)
	err = model.CardValueTokenUpdate(cardToken, req.Card)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responseBase
		response.Message = "Erro ao salvar token"
		response.Code = "96"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := responseBase
	response.Message = "Tokenrização realizadao com sucesso"
	response.Code = "00"
	response.Brand = brand
	response.CardToken = cardToken
	_ = json.NewEncoder(w).Encode(response)
}

func AutorizationCardHanler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req entity.AuthorizationRequest
	response := entity.ValidationResponse{}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(entity.ValidationResponse{
			Message:         "Payload inválido",
			Code:            "96",
			Brand:           "",
			CardToken:       "",
			ExpirationMonth: "",
			ExpirationYear:  "",
			LastFourDigits:  "",
		})
		return
	}

	
	card, err := model.CardValuesByToken(req.Token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "Cartão inexistente na nossa base."
		response.Code = "16"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	if !validation.IsValidLuhn(card.Card) {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "Cartão inválido"
		response.Code = "14"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	resultValidCard, total := validation.ValidationTotal(card.Card, req.Amount)

	if !resultValidCard && total == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "Saldo insuficiente"
		response.Code = "51"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	error := model.CardValueCaptureUpdate(req.Token);

	if error == nil {
		w.WriteHeader(http.StatusOK)
		response.Message = "Autorização realizada com sucesso"
		response.Code = "00"
		_ = json.NewEncoder(w).Encode(response)
		return
	}
}

func CaptureCardHandler(w http.ResponseWriter, r *http.Request) {

	var req entity.ValidationRequest

	result, err := model.CardValuesByCard(req.Card)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(entity.ValidationResponse{
			Message: "Erro ao buscar saldo",
			Code:    "14",
		})
		return
	}

	error := model.CardValuesUpdate(req.Card, req.Total, result.Total)

	if error != nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(entity.ValidationResponse{
			Message: "Erro ao alterar o saldo",
			Code:    "13",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(entity.ValidationResponse{
		Message: "Transacao capturada com sucesso",
		Code:    "01",
	})
}
