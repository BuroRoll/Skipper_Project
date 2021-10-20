package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `json:"email" binding:"required" gorm:"index:unique"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
}
