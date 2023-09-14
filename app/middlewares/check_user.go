package middlewares

import (
	"api-ppob/app/models"
	"api-ppob/app/utils"
	"api-ppob/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		token, err := utils.VerifyTokenHeader(c, secretKey)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  401,
				"message": "Token tidak tidak valid atau kadaluwarsa",
				"data":    nil,
			})
			return
		}

		var count int64

		db := database.Connect()

		email := utils.GetUserFromToken(c)

		db.Model(&models.User{}).Where("email = ? ", email).Count(&count)

		if count == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  101,
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}

		c.Next()

	}
}
