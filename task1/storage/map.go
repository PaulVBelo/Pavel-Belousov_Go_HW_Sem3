package storage

import (
	"strings"
	"task1/book"
	"task1/library"
)

type StorageMap struct {
	books      map[int]book.Book
	GenerateID func() int
}

func NewLibraryMap(generateID func() int) library.Library {
	return &StorageMap{
		books:      make(map[int]book.Book),
		GenerateID: generateID,
	}
}


func (l *StorageMap) AddBook(title, author string) {
	if (author == "") {
		author = "unknown"
	}
	id := l.GenerateID()
	l.books[id] = book.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}

func (l *StorageMap) GetBookByID(id int) (book.Book, bool) {
	b, ok := l.books[id]
	return b, ok
}

func (l *StorageMap) GetBookByTitle(title string) (book.Book, bool) {
	for _, b := range l.books {
		if strings.EqualFold(b.Title, title) {
			return b, true
		}
	}
	return book.Book{}, false
}

func (l *StorageMap) Search(title string) ([]book.Book, bool) {
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

func (l *StorageMap) GetAll() []book.Book {
	var results []book.Book
	for _, b := range l.books {
		results = append(results, b)
	}
	return results
}

func (l *StorageMap) AddInitiatedBook(book book.Book) {
	_, ok := l.books[book.ID]
	if !ok {
		l.books[book.ID] = book
	}
}