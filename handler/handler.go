package handler

import (
	"adquirer/model"
	"adquirer/validation"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type ValidationRequest struct {
	Card            string  `json:"card"`
	CVV             string  `json:"cvv"`
	Venc            string  `json:"venc"`
	HolderName      string  `json:"holderName"`
	ExpirationMonth string  `json:"expirationMonth"`
	ExpirationYear  string  `json:"expirationYear"`
	Total           float64 `json:"total"`
}

type ValidationResponse struct {
	Message         string `json:"message"`
	Code            string `json:"code"`
	Brand           string `json:"brand"`
	CardToken       string `json:"cardToken"`
	ExpirationMonth string `json:"expirationMonth"`
	ExpirationYear  string `json:"expirationYear"`
	LastFourDigits  string `json:"lastFourDigits"`
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

	expirationMonth, expirationYear := extractExpiration(req)
	responseBase := ValidationResponse{
		Brand:           "",
		CardToken:       "",
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		LastFourDigits:  getLastFourDigits(req.Card),
	}

	w.Header().Set("Content-Type", "application/json")
	if !validation.IsValidLuhn(req.Card) {
		w.WriteHeader(http.StatusOK)
		response := responseBase
		response.Message = "Cartão inválido"
		response.Code = "14"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	resultValidCard, total := validation.ValidationTotal(req.Card, req.Total)

	if !resultValidCard && total == 0 {
		w.WriteHeader(http.StatusOK)
		response := responseBase
		response.Message = "Saldo insuficiente"
		response.Code = "51"
		_ = json.NewEncoder(w).Encode(response)
		return
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

	w.WriteHeader(http.StatusOK)
	response := responseBase
	response.Message = "Transacao autorizada com sucesso"
	response.Code = "00"
	response.Brand = brand
	response.CardToken = cardToken
	_ = json.NewEncoder(w).Encode(response)
}

func extractExpiration(req ValidationRequest) (string, string) {
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

func getLastFourDigits(card string) string {
	clean := strings.NewReplacer(" ", "", "-", "").Replace(strings.TrimSpace(card))
	if len(clean) <= 4 {
		return clean
	}

	return clean[len(clean)-4:]
}

func CaptureCardHandler(w http.ResponseWriter, r *http.Request) {

	var req ValidationRequest

	result, err := model.CardValuesByCard(req.Card)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ValidationResponse{
			Message: "Erro ao buscar saldo",
			Code:    "14",
		})
		return
	}

	error := model.CardValuesUpdate(req.Card, req.Total, result.Total)

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
		Message: "Transacao capturada com sucesso",
		Code:    "01",
	})
}
