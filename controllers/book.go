package controllers

import (
	"books-api/models"
	"books-api/repository"
	"books-api/utils"
	"database/sql"
	"encoding/json"
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

// AddBook inserts a book into database
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book) // gets the book from the request

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := repository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text plani")
		utils.SendSuccess(w, bookID)
	}
}

// UpdateBook updates a book
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book) // gets the book from request body

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := repository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

// RemoveBook deletes a book
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		bookRepo := repository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}
	}
}
