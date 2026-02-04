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
	GetByID(id uint) (*Payment, error)
	Update(id uint, payment Payment) error
	Delete(id uint) error
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

func (r *paymentRepository) GetByID(id uint) (*Payment, error) {
	var payment Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(id uint, payment Payment) error {
	return r.db.Model(&Payment{}).Where("id = ?", id).Updates(payment).Error
}

func (r *paymentRepository) Delete(id uint) error {
	return r.db.Delete(&Payment{}, id).Error
}
