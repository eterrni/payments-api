package service

import (
	"errors"

	"github.com/eterrni/payments-api/internal/repository"
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

	return s.repo.CreatePayment(repository.Payment{
		Amount:   payment.Amount,
		Currency: payment.Currency,
	})
}

func (s *PaymentService) GetPayment(id uint) (*repository.Payment, error) {
	return s.repo.GetByID(id)
}

func (s *PaymentService) UpdatePayment(id uint, payment PaymentRequest) error {
	if payment.Amount <= 0 {
		return errors.New("invalid payment amount")
	}
	return s.repo.Update(id, repository.Payment{
		Amount:   payment.Amount,
		Currency: payment.Currency,
	})
}

func (s *PaymentService) DeletePayment(id uint) error {
	return s.repo.Delete(id)
}
