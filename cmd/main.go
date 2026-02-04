package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error
	dsn := os.Getenv("DB_DSN")
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	db.AutoMigrate(&repository.Payment{})
}

func main() {

	r := mux.NewRouter()

}
