package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB //database

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")) //Build connection string
	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "skipper_user", "skipper", "22012011")
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Printf("error %s", err)
	}
	db = conn
	//err = db.Debug().AutoMigrate(&User{}, &Catalog0{}, &Catalog1{}, &Catalog2{}, &Catalog3{}) //Database migration
	err = db.Debug().AutoMigrate(&Catalog0{}) //Database migration
	err = db.Debug().AutoMigrate(&Catalog1{}) //Database migration
	err = db.Debug().AutoMigrate(&Catalog2{}) //Database migration
	err = db.Debug().AutoMigrate(&Catalog3{}) //Database migration
	if err != nil {
		fmt.Println(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
