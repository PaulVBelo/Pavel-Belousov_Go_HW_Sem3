package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Book struct {
	title  string
	author string
	ID     int
}

type Library interface {
	AddBookNA(title string)
	AddBook(title, author string)
	GetBookByID(id int) (Book, bool)
	GetBookByTitle(title string) (Book, bool)
	Search(title string) ([]Book, bool)
	GetAll() []Book
	AddInitiatedBook(book Book)
}

type LibraryImpl struct {
	books      map[int]Book
	generateID func() int
}

type LibraryImpl2 struct {
	books      []Book
	generateID func() int
}

// Lib w/ map

func NewLibrary(generateID func() int) Library {
	return &LibraryImpl{
		books:      make(map[int]Book),
		generateID: generateID,
	}
}

func (l *LibraryImpl) AddBookNA(title string) {
	id := l.generateID()
	l.books[id] = Book{
		ID:     id,
		title:  title,
		author: "unknown",
	}
}

func (l *LibraryImpl) AddBook(title, author string) {
	id := l.generateID()
	l.books[id] = Book{
		ID:     id,
		title:  title,
		author: author,
	}
}

func (l *LibraryImpl) GetBookByID(id int) (Book, bool) {
	b, ok := l.books[id]
	return b, ok
}

func (l *LibraryImpl) GetBookByTitle(title string) (Book, bool) {
	for _, b := range l.books {
		if strings.EqualFold(b.title, title) {
			return b, true
		}
	}
	return Book{}, false
}

func (l *LibraryImpl) Search(title string) ([]Book, bool) {
	var results []Book
	for _, b := range l.books {
		if strings.EqualFold(b.title, title) {
			results = append(results, b)
		}
	}
	if len(results) > 0 {
		return results, true
	} else {
		return []Book{}, false
	}
}

func (l *LibraryImpl) GetAll() []Book {
	var results []Book
	for _, b := range l.books {
		results = append(results, b)
	}
	return results
}

func (l *LibraryImpl) AddInitiatedBook(book Book) {
	_, ok := l.books[book.ID]
	if !ok {
		l.books[book.ID] = book
	}

}

// Lib w/ slices

func NewLibraryWithSlices(generateId func() int) Library {
	return &LibraryImpl2{
		books:      make([]Book, 0),
		generateID: generateId,
	}
}

func (l *LibraryImpl2) AddBookNA(title string) {
	id := l.generateID()
	l.books[id] = Book{
		ID:     id,
		title:  title,
		author: "unknown",
	}
}

func (l *LibraryImpl2) AddBook(title, author string) {
	id := l.generateID()
	l.books[id] = Book{
		ID:     id,
		title:  title,
		author: author,
	}
}

func (l *LibraryImpl2) GetBookByID(id int) (Book, bool) {
	for _, b := range l.books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}

func (l *LibraryImpl2) GetBookByTitle(title string) (Book, bool) {
	for _, b := range l.books {
		if strings.EqualFold(b.title, title) {
			return b, true
		}
	}
	return Book{}, false
}

func (l *LibraryImpl2) Search(title string) ([]Book, bool) {
	var results []Book
	for _, b := range l.books {
		if strings.EqualFold(b.title, title) {
			results = append(results, b)
		}
	}
	if len(results) > 0 {
		return results, true
	} else {
		return []Book{}, false
	}
}

func (l *LibraryImpl2) GetAll() []Book {
	return l.books
}

func (l *LibraryImpl2) AddInitiatedBook(book Book) {
	for _, b := range l.books {
		if b.ID == book.ID {
			return
		}
	}
	l.books = append(l.books, book)
}

// Transfer - Тяжёлая операция перевода одной библиотеки в другую
// Написал отдельно от интерфейса, чтобы его дальше не заполнять
func Transfer(from, to Library) {
	books := from.GetAll()
	for _, b := range books {
		to.AddInitiatedBook(b)
	}
}

// Для удобства

func printResponseSingle(b Book, ok bool) {
	if ok {
		fmt.Printf("Found book: %d %s by %s \n", b.ID, b.title, b.author)
	} else {
		fmt.Println("No books???")
	}
}

func printResponseMulti(books []Book, ok bool) {
	if ok {
		fmt.Print("Found books: ")
		for _, b := range books {
			fmt.Printf("%d %s by %s; ", b.ID, b.title, b.author)
		}
		fmt.Println()
	} else {
		fmt.Println("No books???")
	}
}

func genRandID() int {
	return rand.Int()
}

func genTimebasedId() int {
	return int(time.Now().UnixNano())
}

func main() {
	lib := NewLibrary(genRandID)

	lib.AddBookNA("notebook")
	lib.AddBook("Urban Dynamics", "Jay W. Forrester")
	lib.AddBook("Hobbit, or There and Back Again", "J.R.R. Tolkien")
	printResponseSingle(lib.GetBookByTitle("Urban Dynamics"))

	lib.(*LibraryImpl).generateID = genTimebasedId

	lib.AddBook("Notebook", "Some MIPT student")
	printResponseSingle(lib.GetBookByID(42))
	printResponseMulti(lib.Search("NoTeBOOk"))

	lib2 := NewLibrary(genRandID)
	lib2.AddBook("NotEbooK", "sus amogus")
	printResponseSingle(lib2.GetBookByTitle("Notebook"))

	Transfer(lib, lib2)
	lib.AddBookNA("Goblin Slayer") //Не будет в lib2
	printResponseMulti(lib2.Search("Notebook"))
	printResponseMulti(lib2.GetAll(), true)
}
