package library

import "task1/book"

type Library interface {
	AddBook(title, author string)
	GetBookByID(id int) (book.Book, bool)
	GetBookByTitle(title string) (book.Book, bool)
	Search(title string) ([]book.Book, bool)
	GetAll() []book.Book
	AddInitiatedBook(book book.Book)
}

// Transfer - Тяжёлая операция перевода одной библиотеки в другую
// Написал отдельно от интерфейса, чтобы его дальше не заполнять
func Transfer(from, to Library) {
	books := from.GetAll()
	for _, b := range books {
		to.AddInitiatedBook(b)
	}
}