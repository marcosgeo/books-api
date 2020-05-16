package controllers

import (
	"books-api/models"
	"books-api/repository"
	"books-api/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Controller struct
type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

// GetBook gets a book by id
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)

		bookRepo := repository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}
