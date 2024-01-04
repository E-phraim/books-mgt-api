package config

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbConn, err := gorm.Open("mysql", "root:administration@tcp(localhost:3306)/booksmgt?charset=utf8&parseTime=True&loc=Local")
	// dbConn, err := gorm.Open("mysql", "root:administration/booksmgt?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = dbConn
}

func GetDB() *gorm.DB {
	return db
}
