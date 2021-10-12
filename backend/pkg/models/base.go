package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func init() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "danilkonkov", "skipper", "") //Build connection string

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Printf("error %s", err)
	}
	db = conn
	err = db.Debug().AutoMigrate(&User{}) //Database migration
	if err != nil {
		fmt.Println(err)
	}
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
