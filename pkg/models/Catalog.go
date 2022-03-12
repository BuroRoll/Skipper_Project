package models

import "gorm.io/gorm"

type Catalog0 struct {
	gorm.Model
	Name   string     `json:"name0"`
	Child0 []Catalog1 `gorm:"ForeignKey:ParentId"`
}

type Catalog1 struct {
	gorm.Model
	ParentId uint
	Name     string     `json:"name1"`
	Child1   []Catalog2 `gorm:"ForeignKey:ParentId"`
}

type Catalog2 struct {
	gorm.Model
	ParentId uint
	Name     string     `json:"name2"`
	Child2   []Catalog3 `gorm:"ForeignKey:ParentId"`
}

type Catalog3 struct {
	gorm.Model
	ParentId uint
	Name3    string `json:"name3"`
	Count    uint   `json:"count"`
}
