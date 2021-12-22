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
	ClassBooking   []*Class        `gorm:"many2many:user_classbooking;"`
	//Mentor information
	Specialization   string
	Educations       []Education        `gorm:"ForeignKey:ParentId"`
	WorkExperiences  []WorkExperience   `gorm:"ForeignKey:ParentId"`
	Classes          []Class            `gorm:"ForeignKey:ParentId"`
	OtherInformation []OtherInformation `gorm:"ForeignKey:ParentId"`
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

type Education struct {
	gorm.Model
	ParentId    uint
	Institution string
	StartYear   int
	EndYear     int
	Degree      string
}

type WorkExperience struct {
	gorm.Model
	ParentId     uint
	Organization string
	StartYear    int
	EndYear      int
}

type OtherInformation struct {
	gorm.Model
	ParentId uint
	Data     string
}

type Pagination struct {
	Limit  int      `json:"limit"`
	Page   int      `json:"page"`
	Sort   string   `json:"sort"`
	Search []string `json:"search"`
}
