package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/octodemo/advanced-security-go/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error

	os.Remove("./bookstore.db")

	models.DB, err = sql.Open("sqlite3", "./bookstore.db")
	if err != nil {
		log.Fatal(err)
	}
	defer models.DB.Close()

	sqlStmt := `
	CREATE TABLE books (
		name varchar(255) NOT NULL,
		author varchar(255) NOT NULL,
		read varchar(255) NOT NULL
	);

	INSERT INTO books (name, author, read) VALUES
	("The Hobbit", "JRR Tolkien", "True"),
	("The Fellowship of the Ring", "JRR Tolkien", "True"),
	("The Eye of the World", "Robert Jordan", "False"),
	("A Game of Thrones", "George R. R. Martin", "True"),
	("The Way of Kings", "Brandon Sanderson", "False");
	`
	_, err = models.DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	_, err = models.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	author := r.URL.Query().Get("author")
	read := r.URL.Query().Get("read")

	if len(name) > 0 {
		bks, err := models.NameQuery(name)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s\n", bk.Title, bk.Author, bk.Read)
		}

	} else if len(author) > 0 {
		bks, err := models.AuthorQuery(author)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s\n", bk.Title, bk.Author, bk.Read)
		}

	} else if len(read) > 0 {
		bks, err := models.ReadQuery(read)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s\n", bk.Title, bk.Author, bk.Read)
		}

	} else {
		bks, err := models.AllBooks()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s\n", bk.Title, bk.Author, bk.Read)
		}
	}
}
