package main

import "database/sql"

// Book structure to be serialized
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": Book{Title: "Advanced Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

// AllBooks returns a slice of all books
func AllBooks(db *sql.DB) []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// GetBook returns the book for a given ISBN
func GetBook(db *sql.DB, isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}
