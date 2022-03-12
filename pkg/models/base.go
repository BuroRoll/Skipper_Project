package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var db *gorm.DB

func init() {
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../..")
	err := godotenv.Load(Root + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbUri := os.Getenv("DATABASE_URL")
	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatalf("error %s", err)
	}
	db = conn
	err = db.AutoMigrate(
		&User{},
		&Catalog0{},
		&Catalog1{},
		&Catalog2{},
		&Catalog3{},
		&Messenger{},
		&Communication{},
		&Education{},
		&WorkExperience{},
		&Class{},
		&TheoreticClass{},
		&PracticClass{},
		&KeyClass{},
		&OtherInformation{},
	)
	err = db.AutoMigrate(&BookingTime{})
	err = db.AutoMigrate(&UserClass{})
	err = db.SetupJoinTable(&User{}, "ClassBooking", &UserClass{})

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
