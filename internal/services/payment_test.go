package service

import (
	"errors"
	"testing"

	"github.com/eterrni/payments-api/internal/repository"
)

type mockPaymentRepository struct {
	createErr error
	getResult *repository.Payment
	getErr    error
	updateErr error
	deleteErr error
}

func (m *mockPaymentRepository) CreatePayment(payment repository.Payment) error {
	return m.createErr
}

func (m *mockPaymentRepository) GetByID(id uint) (*repository.Payment, error) {
	return m.getResult, m.getErr
}

func (m *mockPaymentRepository) Update(id uint, payment repository.Payment) error {
	return m.updateErr
}

func (m *mockPaymentRepository) Delete(id uint) error {
	return m.deleteErr
}

func TestPaymentService_CreatePayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockPaymentRepository{}
		svc := NewPaymentService(repo)

		err := svc.CreatePayment(PaymentRequest{Amount: 100.5, Currency: "USD"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if repo.createErr != nil {
			t.Fatal("createErr should be nil")
		}
	})

	t.Run("invalid amount zero", func(t *testing.T) {
		repo := &mockPaymentRepository{}
		svc := NewPaymentService(repo)

		err := svc.CreatePayment(PaymentRequest{Amount: 0, Currency: "USD"})
		if err == nil {
			t.Fatal("expected error for zero amount")
		}
		if err.Error() != "invalid payment amount" {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid amount negative", func(t *testing.T) {
		repo := &mockPaymentRepository{}
		svc := NewPaymentService(repo)

		err := svc.CreatePayment(PaymentRequest{Amount: -10, Currency: "USD"})
		if err == nil {
			t.Fatal("expected error for negative amount")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &mockPaymentRepository{createErr: errors.New("db error")}
		svc := NewPaymentService(repo)

		err := svc.CreatePayment(PaymentRequest{Amount: 100, Currency: "USD"})
		if err == nil {
			t.Fatal("expected error from repository")
		}
		if err.Error() != "db error" {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestPaymentService_GetPayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expected := &repository.Payment{ID: 1, Amount: 50, Currency: "EUR"}
		repo := &mockPaymentRepository{getResult: expected}
		svc := NewPaymentService(repo)

		payment, err := svc.GetPayment(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if payment != expected {
			t.Errorf("got %+v, want %+v", payment, expected)
		}
	})

	t.Run("not found", func(t *testing.T) {
		repo := &mockPaymentRepository{getErr: errors.New("record not found")}
		svc := NewPaymentService(repo)

		_, err := svc.GetPayment(999)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestPaymentService_UpdatePayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockPaymentRepository{}
		svc := NewPaymentService(repo)

		err := svc.UpdatePayment(1, PaymentRequest{Amount: 200, Currency: "USD"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid amount", func(t *testing.T) {
		repo := &mockPaymentRepository{}
		svc := NewPaymentService(repo)

		err := svc.UpdatePayment(1, PaymentRequest{Amount: 0, Currency: "USD"})
		if err == nil {
			t.Fatal("expected error for invalid amount")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &mockPaymentRepository{updateErr: errors.New("update failed")}
		svc := NewPaymentService(repo)

		err := svc.UpdatePayment(1, PaymentRequest{Amount: 100, Currency: "USD"})
		if err == nil {
			t.Fatal("expected error from repository")
		}
	})
}

func TestPaymentService_DeletePayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockPaymentRepository{}
		svc := NewPaymentService(repo)

		err := svc.DeletePayment(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &mockPaymentRepository{deleteErr: errors.New("delete failed")}
		svc := NewPaymentService(repo)

		err := svc.DeletePayment(1)
		if err == nil {
			t.Fatal("expected error from repository")
		}
	})
}
