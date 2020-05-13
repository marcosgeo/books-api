package main

import (
	"books-api/driver"
	"books-api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db = driver.ConnectDB()

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	driver.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []models.Book{}

	rows, err := db.Query("select id, title, author, year from books;")
	logFatal(err)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var book models.Book
	rows := db.QueryRow("select id, title, author, year from books where id= $1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
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
