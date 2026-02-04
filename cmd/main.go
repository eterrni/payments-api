package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error
	dsn := os.Getenv("DB_DSN")
	gorm.Open()
}

func main() {

}
