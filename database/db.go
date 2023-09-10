package database

import (
	"api-ppob/app/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	DB_HOST := viper.GetString("database.host")
	DB_PORT := viper.GetString("database.port")
	DB_USER := viper.GetString("database.user")
	DB_PASSWORD := viper.GetString("database.password")
	DB_NAME := viper.GetString("database.database")
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
