package main

import (
	"books-api/controllers"
	"books-api/driver"
	"books-api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow(`
		insert into books (title, author, year) values ($1, $2, $3) returning id;`,
		book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec(
		`update books set title=$1, author=$2, year=$3 where id=$4 returning id`,
		&book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get a map of id and values from each request param

	id, _ := strconv.Atoi(params["id"])

	result, err := db.Exec("delete from books where id = $1", id)
	rowsAffected, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsAffected)
}
