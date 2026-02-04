package service

import (
	"errors"

	"github.com/payment-api/internal/repository"
)

type PaymentService struct {
	repo repository.PaymentRepository
}

type PaymentRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(payment PaymentRequest) error {
	if payment.Amount <= 0 {
		return errors.New("invalid payment amount")
	}

	return s.repo.CreatePayment(payment)
}
