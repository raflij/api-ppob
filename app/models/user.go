package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_Name string `gorm:"size:255;not null" validate:"required" json:"first_name"`
	Last_Name  string `gorm:"size:255;not null" validate:"required" json:"last_name"`
	Email      string `gorm:"size:255;unique;not null" json:"email"`
	Password   string `gorm:"size:255;not null" json:"password"`
}
