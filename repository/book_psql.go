package repository

import (
	"books-api/models"
	"database/sql"
	"log"
)

// BookRepository represents a book in the database
type BookRepository struct{}

// GetBooks returns a list o books
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select id, title, author, year from books")
	if err != nil {
		logFatal(err)
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)

	}
	if err != nil {
		logFatal(err)
		return []models.Book{}, err
	}
	return books, nil
}

// GetBook gets a book by id
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	rows := db.QueryRow("select id, title, author, year from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	return book, err
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}
