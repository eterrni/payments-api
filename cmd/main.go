package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eterrni/payments-api/internal/handlers"
	"github.com/eterrni/payments-api/internal/repository"
	"github.com/eterrni/payments-api/internal/service"
	"github.com/eterrni/payments-api/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	// Настройка логирования
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Подключение к базе данных
	var err error
	dsn := os.Getenv("DB_DSN")
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Автоматическая миграция
	db.AutoMigrate(&repository.Payment{})
}

func main() {
	r := mux.NewRouter()

	// Миддлвары
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoveryMiddleware)

	// Хендлеры
	ph := handlers.NewPaymentHandler(service.NewPaymentService(repository.NewPaymentRepository(db)))
	r.HandleFunc("/payments", ph.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", ph.GetPayment).Methods("GET")
	r.HandleFunc("/payments/{id}", ph.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", ph.DeletePayment).Methods("DELETE")

	// Запуск сервера
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
