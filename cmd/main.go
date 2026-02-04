package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/eterrni/payments-api/internal/handlers"
	"github.com/eterrni/payments-api/internal/repository"
	service "github.com/eterrni/payments-api/internal/services"
	"github.com/eterrni/payments-api/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error
	dsn := os.Getenv("DB_DSN")
	if dsn != "" && !strings.Contains(dsn, "sslmode=") {
		dsn = strings.TrimSpace(dsn) + " sslmode=disable"
	}
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.AutoMigrate(&repository.Payment{})
}

func main() {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoveryMiddleware)

	paymentSvc := service.NewPaymentService(repository.NewPaymentRepository(db))
	ph := handlers.NewPaymentHandler(&paymentSvc)
	r.HandleFunc("/payments", ph.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", ph.GetPayment).Methods("GET")
	r.HandleFunc("/payments/{id}", ph.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", ph.DeletePayment).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
