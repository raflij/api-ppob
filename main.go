package main

import (
	"api-ppob/database"
	"api-ppob/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	routers.InitRoutes(r)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	database.Init()
	database.Connect()
	r.Run(":" + viper.GetString("appPort"))

}
