package models

import (
	"fmt"

	"github.com/e-phraim/books-mgt-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

	fmt.Println("DB Mig. Suc.")
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	fmt.Println("book created")
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetOneBook(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteOneBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
