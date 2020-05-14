package controllers

import (
	"books-api/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Controller struct
type Controller struct{}

var books []models.Book

// GetBooks ...
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
