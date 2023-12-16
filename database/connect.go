package database

import (
	"log"
	"os"

	"github.com/TomGaleano/Golang/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loagding .env file.")
	}
	dsn := os.Getenv("DSN")
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database.")
	} else {
		log.Println("Connected succesfully.")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
	)
}
