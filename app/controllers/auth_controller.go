package controllers

import (
	"api-ppob/app/models"
	"api-ppob/app/utils"
	"api-ppob/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email" validate:"required,email,max=96"`
	Password string `json:"password" validate:"required,max=48"`
}

type Register struct {
	First_Name string `json:"first_name" validate:"required,max=64"`
	Last_Name  string `json:"last_name" validate:"required,max=64"`
	Email      string `json:"email" validate:"required,email,max=96"`
	Password   string `json:"password" validate:"required,min=8,eqfield=RePassword"`
	RePassword string `json:"repassword" validate:"required,min=8,max=48"`
}

func RegisterUser(c *gin.Context) {
	var dataInput Register

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(400, gin.H{
			"status":  101,
			"message": "invalid parameters",
			"data":    err.Error(),
		})
		return
	}

	checkValidation, statusValidate := utils.ValidationCustom(c, dataInput)

	if statusValidate {
		c.JSON(400, gin.H{
			"status":  101,
			"message": checkValidation,
			"data":    nil,
		})
		return
	}

	var db = database.Connect()

	var user models.User

	userCheck := db.Where(models.User{
		Email: dataInput.Email,
	}).First(&user)

	if userCheck.RowsAffected > 0 {
		c.JSON(400, gin.H{
			"status":  102,
			"message": "Email sudah terdaftar",
			"data":    nil,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dataInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  103,
			"message": "Gagal membuat akun",
			"data":    nil,
		})
		return
	}

	db.Create(&models.User{
		First_Name: dataInput.First_Name,
		Last_Name:  dataInput.Last_Name,
		Email:      dataInput.Email,
		Password:   string(hashedPassword),
	})

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Akun berhasil dibuat",
		"data": gin.H{
			"first_name": dataInput.First_Name,
			"last_name":  dataInput.Last_Name,
			"email":      dataInput.Email,
			"password":   dataInput.Password,
		},
	})

}

func LoginUser(c *gin.Context) {
	var dataInput Login

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(400, gin.H{
			"status":  101,
			"message": "invalid parameters",
			"data":    err.Error(),
		})
		return
	}

	checkValidation, statusValidate := utils.ValidationCustom(c, dataInput)

	if statusValidate {
		c.JSON(400, gin.H{
			"status":  101,
			"message": checkValidation,
			"data":    nil,
		})
		return
	}

	var db = database.Connect()

	var user models.User

	if err := db.Where(&models.User{
		Email: dataInput.Email,
	}).First(&user).Error; err != nil {
		c.JSON(401, gin.H{
			"status":  102,
			"message": "Email atau password salah",
			"data":    nil,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataInput.Password)); err != nil {
		c.JSON(401, gin.H{
			"status":  103,
			"message": "Email atau password salah",
			"data":    nil,
		})
		return
	}

	dataClaims := map[string]interface{}{
		"id":    int(user.ID),
		"email": user.Email,
	}

	token, err := utils.Sign(dataClaims, viper.GetString("secretKey.jwt"), 3600)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  104,
			"message": "Gagal membuat access_token",
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Login berhasil",
		"data": gin.H{
			"token": token,
		},
	})

	return
}
