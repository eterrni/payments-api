package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/payment-api/internal/service"
	"github.com/payment-api/pkg/utils"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment service.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreatePayment(payment); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not create payment")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Payment created"})
}

func (h *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
}

func (h *PaymentHandler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
}

func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {
}
