package main

import (
	"database/sql"
)

func createTable(db *sql.DB) sql.Result {
	query := "create table books (id serial, title varchar, author varchar, year varchar);"

	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return result
}

func insertData(db *sql.DB) sql.Result {
	query := `
		insert into books (title, author, year) values ('Golang is great', 'Mr. Great', '2012');
		insert into books (title, author, year) values ('Golang is magnific', 'Mr. Magnific', '2013');
		insert into books (title, author, year) values ('Os sertões', 'Euclides da Cunha', '1895');
		insert into books (title, author, year) values ('Memórias póstumas de Bráz Cubas', 'Machado de Assis', '1878');
		insert into books (title, author, year) values ('O ultimo dos justos', 'Andre Schwart-Bart', '1960');
	`

	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return result
}
