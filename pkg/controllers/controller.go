package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/e-phraim/books-mgt-api/pkg/models"
	"github.com/e-phraim/books-mgt-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOneBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}

	bookDetails, _ := models.GetOneBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	NewEntry := &models.Book{}
	utils.ParseBody(r, NewEntry)

	b := NewEntry.CreateBook()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println(NewEntry)
}

func DeleteOneBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteOneBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateOneBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}

	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetOneBook(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
