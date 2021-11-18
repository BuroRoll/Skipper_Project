package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//Base information
	Email          string
	Phone          string `gorm:"index:unique"`
	Password       string
	FirstName      string
	SecondName     string
	Patronymic     string
	DateOfBirthday string
	Description    string
	Time           string
	ProfilePicture string          `gorm:"default:'default_profile_picture.jpeg'"`
	IsMentor       bool            `gorm:"default:false"`
	IsVerifyEmail  bool            `gorm:"default:false"`
	IsVerifyPhone  bool            `gorm:"default:false"`
	Communications []Communication `gorm:"ForeignKey:ParentId"`
	//Mentor information
	Specialization string
}

type Communication struct {
	gorm.Model
	ParentId  uint
	Messenger []*Messenger `gorm:"many2many:messenger_communication;"`
	Login     string
}

type Messenger struct {
	gorm.Model
	Name           string
	Communications []*Communication `gorm:"many2many:messenger_communication;"`
}
