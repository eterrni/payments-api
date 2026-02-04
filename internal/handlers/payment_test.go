package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eterrni/payments-api/internal/repository"
	service "github.com/eterrni/payments-api/internal/services"
	"github.com/gorilla/mux"
)

type mockPaymentService struct {
	createErr error
	getResult *repository.Payment
	getErr    error
	updateErr error
	deleteErr error
}

func (m *mockPaymentService) CreatePayment(payment service.PaymentRequest) error {
	return m.createErr
}

func (m *mockPaymentService) GetPayment(id uint) (*repository.Payment, error) {
	return m.getResult, m.getErr
}

func (m *mockPaymentService) UpdatePayment(id uint, payment service.PaymentRequest) error {
	return m.updateErr
}

func (m *mockPaymentService) DeletePayment(id uint) error {
	return m.deleteErr
}

func TestPaymentHandler_CreatePayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := &mockPaymentService{}
		h := NewPaymentHandler(mock)
		body := map[string]interface{}{"amount": 100.5, "currency": "USD"}
		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.CreatePayment(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("got status %d, want %d", w.Code, http.StatusCreated)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := NewPaymentHandler(&mockPaymentService{})

		req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader([]byte("invalid")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.CreatePayment(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mock := &mockPaymentService{createErr: errors.New("db error")}
		h := NewPaymentHandler(mock)
		body := map[string]interface{}{"amount": 100, "currency": "USD"}
		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		h.CreatePayment(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("got status %d, want %d", w.Code, http.StatusInternalServerError)
		}
	})
}

func TestPaymentHandler_GetPayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expected := &repository.Payment{ID: 1, Amount: 50, Currency: "EUR"}
		mock := &mockPaymentService{getResult: expected}
		h := NewPaymentHandler(mock)

		req := httptest.NewRequest(http.MethodGet, "/payments/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		h.GetPayment(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		h := NewPaymentHandler(&mockPaymentService{})

		req := httptest.NewRequest(http.MethodGet, "/payments/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})
		w := httptest.NewRecorder()

		h.GetPayment(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mock := &mockPaymentService{getErr: errors.New("not found")}
		h := NewPaymentHandler(mock)

		req := httptest.NewRequest(http.MethodGet, "/payments/999", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "999"})
		w := httptest.NewRecorder()

		h.GetPayment(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("got status %d, want %d", w.Code, http.StatusNotFound)
		}
	})
}

func TestPaymentHandler_UpdatePayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := &mockPaymentService{}
		h := NewPaymentHandler(mock)
		body := map[string]interface{}{"amount": 200, "currency": "USD"}
		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPut, "/payments/1", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		h.UpdatePayment(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		h := NewPaymentHandler(&mockPaymentService{})

		req := httptest.NewRequest(http.MethodPut, "/payments/x", bytes.NewReader([]byte("{}")))
		req = mux.SetURLVars(req, map[string]string{"id": "x"})
		w := httptest.NewRecorder()

		h.UpdatePayment(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mock := &mockPaymentService{updateErr: errors.New("update failed")}
		h := NewPaymentHandler(mock)
		body := map[string]interface{}{"amount": 100, "currency": "USD"}
		jsonBody, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPut, "/payments/1", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		h.UpdatePayment(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("got status %d, want %d", w.Code, http.StatusInternalServerError)
		}
	})
}

func TestPaymentHandler_DeletePayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := &mockPaymentService{}
		h := NewPaymentHandler(mock)

		req := httptest.NewRequest(http.MethodDelete, "/payments/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		h.DeletePayment(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		h := NewPaymentHandler(&mockPaymentService{})

		req := httptest.NewRequest(http.MethodDelete, "/payments/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})
		w := httptest.NewRecorder()

		h.DeletePayment(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("service error", func(t *testing.T) {
		mock := &mockPaymentService{deleteErr: errors.New("delete failed")}
		h := NewPaymentHandler(mock)

		req := httptest.NewRequest(http.MethodDelete, "/payments/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		h.DeletePayment(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("got status %d, want %d", w.Code, http.StatusInternalServerError)
		}
	})
}

func TestGetIDFromRequest(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/payments/42", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "42"})

		id, err := getIDFromRequest(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != 42 {
			t.Errorf("got id %d, want 42", id)
		}
	})

	t.Run("missing id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/payments/", nil)
		req = mux.SetURLVars(req, map[string]string{})

		_, err := getIDFromRequest(req)
		if err == nil {
			t.Fatal("expected error for missing id")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/payments/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})

		_, err := getIDFromRequest(req)
		if err == nil {
			t.Fatal("expected error for invalid id")
		}
	})
}
