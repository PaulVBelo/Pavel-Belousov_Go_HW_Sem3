package storage

import (
	"strings"
	"task1/book"
	"task1/library"
)

type StorageSlice struct {
	books      []book.Book
	GenerateID func() int
}

func NewLibrarySlice(generateId func() int) library.Library {
	return &StorageSlice{
		books:      make([]book.Book, 0),
		GenerateID: generateId,
	}
}

func (l *StorageSlice) AddBook(title, author string) {
	id := l.GenerateID()
	if (author == "") {
		author = "unknown"
	}
	l.books = append(l.books, book.Book{
		ID:     id,
		Title:  title,
		Author: author,
	})
}

func (l *StorageSlice) GetBookByID(id int) (book.Book, bool) {
	for _, b := range l.books {
		if b.ID == id {
			return b, true
		}
	}
	return book.Book{}, false
}

func (l *StorageSlice) GetBookByTitle(title string) (book.Book, bool) {
	for _, b := range l.books {
		if strings.EqualFold(b.Title, title) {
			return b, true
		}
	}
	return book.Book{}, false
}

func (l *StorageSlice) Search(title string) ([]book.Book, bool) {
	var results []book.Book
	for _, b := range l.books {
		if strings.EqualFold(b.Title, title) {
			results = append(results, b)
		}
	}
	if len(results) > 0 {
		return results, true
	}
	return nil, false
}

func (l *StorageSlice) GetAll() []book.Book {
	return l.books
}

func (l *StorageSlice) AddInitiatedBook(book book.Book) {
	for _, b := range l.books {
		if b.ID == book.ID {
			return
		}
	}
	l.books = append(l.books, book)
}
