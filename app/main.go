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
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")

	if user == "" || password == "" || db_name == "" {
		log.Fatal("Envoirment variable missing!!")
	}

	url := "postgres://" + user + ":" + password + "@db/" + db_name

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Opening Database: ", err)
	}
	db.AutoMigrate(&Post{})
	return db
}
