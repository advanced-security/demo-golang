package models

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

type Book struct {
	Title  string
	Author string
	Read   string
}

// Get all books in the books table.
func AllBooks() ([]Book, error) {
	query := "SELECT * FROM books"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// Query for books by name. This function contains a SQL Injection issue.
// The user input is not parameterized. Instead of using fmt.Sprintf() to build
// the query, you should be using a parameterized query.
func NameQuery(r string) ([]Book, error) {
	// Fix: rows, err := DB.Query("SELECT * FROM books WHERE name = ?", r)
	rows, err := DB.Query(fmt.Sprintf("SELECT * FROM books WHERE name = '%s'", r))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// Query for books by Author. This function contains a SQL Injection issue.
// The user input is not parameterized. Instead of using fmt.Sprintf() to build
// the query, you should be using a parameterized query.
func AuthorQuery(r string) ([]Book, error) {
	// Fix: rows, err := DB.Query("SELECT * FROM books WHERE author = ?", r)
	rows, err := DB.Query(fmt.Sprintf("SELECT * FROM books WHERE author = '%s'", r))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// Query for books by read.  This function contains a SQL Injection issue.
// The user input is not parameterized. Instead of using fmt.Sprintf() to build
// the query, you should be using a parameterized query.
func ReadQuery(r string) ([]Book, error) {
	// Fix: rows, err := DB.Query("SELECT * FROM books WHERE read = ?", r)
	rows, err := DB.Query(fmt.Sprintf("SELECT * FROM books WHERE read = '%s'", r))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// A helper function to cast the query results to a slice
func makeBookSlice(r *sql.Rows) ([]Book, error) {
	var bks []Book

	for r.Next() {
		var bk Book

		err := r.Scan(&bk.Title, &bk.Author, &bk.Read)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err := r.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
