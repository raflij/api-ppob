package main

import (
	"api-ppob/database"
	"api-ppob/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	routers.InitRoutes(r)

	database.Init()
	database.Connect()
	r.Run(":" + os.Getenv("APP_PORT"))

}
