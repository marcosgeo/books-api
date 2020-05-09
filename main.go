package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book represents a book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()
	books = append(books,
		Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutines", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"},
	)
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Server running...")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get book is called")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book is called")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update book is called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book is called")
}
