package main

import (
	"database/sql"
)

func createTable(db *sql.DB) sql.Result {
	query := "create table if not exists books (id serial, title varchar, author varchar, year varchar);"

	result, err := db.Exec(query)
	logFatal(err) // implemented in main.go
	return result
}

func insertData(db *sql.DB) bool {
	query := `
		insert into books (title, author, year) values ('Golang is great', 'Mr. Great', '2012');
		insert into books (title, author, year) values ('Golang is magnific', 'Mr. Magnific', '2013');
		insert into books (title, author, year) values ('Os sertões', 'Euclides da Cunha', '1895');
		insert into books (title, author, year) values ('Memórias póstumas de Bráz Cubas', 'Machado de Assis', '1878');
		insert into books (title, author, year) values ('O ultimo dos justos', 'Andre Schwart-Bart', '1960');
	`
	rows, _ := db.Query("select count(*) from books;")
	defer rows.Close()
	var count int
	// count the number of rows
	for rows.Next() {
		rows.Scan(&count)
	}

	if count == 0 {
		_, err := db.Exec(query)
		logFatal(err) // implemented in main.go
	}
	return true
}
