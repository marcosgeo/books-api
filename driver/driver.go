package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" //only the driver is usede. none of the package's exported names are visible.
)

var db *sql.DB

// ConnectDB connects to a database
func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	logFatal(err)

	createTable()
	insertData()

	return db
}

func createTable() sql.Result {
	query := "create table if not exists books (id serial, title varchar, author varchar, year varchar);"

	result, err := db.Exec(query)
	logFatal(err) // implemented in main.go
	return result
}

func insertData() bool {
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

func logFatal(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
