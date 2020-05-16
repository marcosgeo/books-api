package controllers

import (
	"books-api/models"
	"books-api/repository"
	"books-api/utils"
	"database/sql"
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
		var error models.Error

		books = []models.Book{}
		bookRepo := repository.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)
		if err != nil {
			logFatal(err)
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Conten-type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
