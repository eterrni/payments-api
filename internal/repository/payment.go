package repository

import (
	"github.com/jinzhu/gorm"
)

type Payment struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type PaymentRepository interface {
	CreatePayment(payment Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) CreatePayment(payment Payment) error {
	return r.db.Create(&payment).Error
}
