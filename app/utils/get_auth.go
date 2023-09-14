package utils

import (
	"api-ppob/app/models"
	"api-ppob/database"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (user models.User) {
	db := database.Connect()

	email := GetUserFromToken(c)
	db.Where(models.User{
		Email: email,
	}).Select("ID, first_name, last_name, email").First(&user)

	return user
}
