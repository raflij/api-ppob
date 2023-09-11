package database

import (
	"api-ppob/app/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.User{},
	)

}

func Connect() *gorm.DB {
	return db
}
