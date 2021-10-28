package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string
	Phone          string `gorm:"index:unique"`
	Password       string
	FirstName      string
	SecondName     string
	Specialization string
	Description    string
	Time           string
	ProfilePicture string `gorm:"default:'default_profile_picture.jpeg'"`
	IsMentor       bool   `gorm:"default:false"`
}
