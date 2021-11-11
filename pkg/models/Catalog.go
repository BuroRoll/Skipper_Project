package models

import "gorm.io/gorm"

//type Catalog struct {
//	gorm.Model
//	Name     string
//	ParentId *uint
//}

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

//type User1 struct {
//	gorm.Model
//	MemberNumber string
//	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number     string
//	UserNumber string
//}

//type Catalog struct {
//	gorm.Model
//	Title string `json:"title"`
//	Id    uint   `json:"index_catalog"`
//	Items []struct {
//		Index int    `json:"index1"`
//		Logo  int    `json:"logo"`
//		Name  string `json:"name"`
//		Items []struct {
//			Index int    `json:"index2"`
//			Title string `json:"title"`
//			Items []struct {
//				Index int    `json:"index3"`
//				Name  string `json:"name"`
//				Count int    `json:"count"`
//			} `json:"items3"`
//		} `json:"items2"`
//	} `json:"items1"`
//}
