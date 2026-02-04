package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eterrni/payments-api/internal/repository"
	service "github.com/eterrni/payments-api/internal/services"
	"github.com/eterrni/payments-api/pkg/utils"
	"github.com/gorilla/mux"
)

type paymentService interface {
	CreatePayment(service.PaymentRequest) error
	GetPayment(uint) (*repository.Payment, error)
	UpdatePayment(uint, service.PaymentRequest) error
	DeletePayment(uint) error
}

type PaymentHandler struct {
	service paymentService
}

func NewPaymentHandler(svc paymentService) *PaymentHandler {
	return &PaymentHandler{service: svc}
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
	id, err := getIDFromRequest(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid payment ID")
		return
	}

	payment, err := h.service.GetPayment(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Payment not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, payment)
}

func (h *PaymentHandler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid payment ID")
		return
	}

	var payment service.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.UpdatePayment(id, payment); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not update payment")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Payment updated"})
}

func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid payment ID")
		return
	}

	if err := h.service.DeletePayment(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not delete payment")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Payment deleted"})
}

func getIDFromRequest(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, strconv.ErrSyntax
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
