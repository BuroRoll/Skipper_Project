package models

import "gorm.io/gorm"

type UserClass struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ClassID   uint `gorm:"primaryKey"`
	ClassType string
	Status    string `gorm:"default:'consideration'"`
	MentiId   uint

	Duration15   bool
	Price15      uint
	Duration30_1 bool
	Price30_1    uint
	Duration30_3 bool
	Price30_3    uint
	Duration30_5 bool
	Price30_5    uint

	Duration60_1 bool
	Price60_1    uint
	Duration60_3 bool
	Price60_3    uint
	Duration60_5 bool
	Price60_5    uint

	Duration90_1 bool
	Price90_1    uint
	Duration90_3 bool
	Price90_3    uint
	Duration90_5 bool
	Price90_5    uint

	Time []BookingTime `gorm:"foreignKey:BookingClassID;references:ID"`

	Communication uint
}

type BookingTime struct {
	gorm.Model
	BookingClassID uint
	Time           string
}