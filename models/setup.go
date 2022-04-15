package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func SetupDB() *gorm.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file or you may forgot to add .env file")
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")
	DBDRIVER := os.Getenv("DB_DRIVER")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err = gorm.Open(DBDRIVER, URL)

	if err != nil {
		panic(err.Error())
	}
	return db
}
