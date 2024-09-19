package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres dbname=goapi port=5432"
var DB *gorm.DB

func DatabaseConnection() {
	var error error

	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
