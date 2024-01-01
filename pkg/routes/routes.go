package routes

import (
	"github.com/e-phraim/books-mgt-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var BooksMgtRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.GetOneBook).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.UpdateOneBook).Methods("PUT")
	r.HandleFunc("/book/{bookID}", controllers.DeleteOneBook).Methods("DELETE")

}
