package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	dbConn, err := gorm.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}
	db = dbConn
}

func GetDB() *gorm.DB {
	return db
}
