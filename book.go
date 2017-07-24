package main

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"1": Book{Title: "pierwsza", Author: "pierwszy autor", ISBN: "isbn1", Description: "opeis1"},
	"2": Book{Title: "druga", Author: "drugi autor", ISBN: "isbn2"},
}

func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}
