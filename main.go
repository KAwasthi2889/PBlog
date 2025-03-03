package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = OpenDB()
	TUI()
}

func OpenDB() *gorm.DB {
	dbu := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbu), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Opening Database: ", err)
	}
	db.AutoMigrate(&Post{})
	return db
}
