package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	SenderId     *uint
	RecipienId   uint
	Text         string
	Rating       uint
	Anonymous    bool
	LessonsCount *uint
}
