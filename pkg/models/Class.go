package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	ParentId    uint
	ClassName   string
	Description string
	Tags        []*Catalog3 `gorm:"many2many:catalog3_class;"`

	TheoreticClass TheoreticClass `gorm:"foreignkey:ClassParentId; default:null'; OnDelete:CASCADE;"`
	PracticClass   PracticClass   `gorm:"foreignkey:ClassParentId; default:null'; OnDelete:CASCADE;"`
	KeyClass       KeyClass       `gorm:"foreignkey:ClassParentId; default:null'; OnDelete:CASCADE;"`
}

type TheoreticClass struct {
	gorm.Model
	ClassParentId uint
	Duration15    bool
	Price15       uint

	Duration30 bool
	Price30    uint

	Duration60 bool
	Price60    uint

	Duration90 bool
	Price90    uint

	Time string
}

type PracticClass struct {
	gorm.Model
	ClassParentId uint
	Duration15    bool
	Price15       uint

	Duration30 bool
	Price30    uint

	Duration60 bool
	Price60    uint

	Duration90 bool
	Price90    uint

	Time string
}

type KeyClass struct {
	gorm.Model
	ClassParentId uint
	Duration15    bool
	Price15       uint

	FullTime      bool
	PriceFullTime uint

	Time string
}
