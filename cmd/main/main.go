package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/e-phraim/books-mgt-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BooksMgtRoutes(r)

	http.Handle("/", r)

	fmt.Printf("Server at port http://localhost:8081\n")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
