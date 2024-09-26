package main

import (
	"fmt"
	"math/rand"
	"task1/book"
	"task1/library"
	"task1/storage"
	"time"
)

// Для удобства

func printResponseSingle(b book.Book, ok bool) {
	if ok {
		fmt.Printf("Found book: %d %s by %s \n", b.ID, b.Title, b.Author)
		return
	}
	fmt.Println("No books???")
}

func printResponseMulti(books []book.Book, ok bool) {
	if ok {
		fmt.Print("Found books: ")
		for _, b := range books {
			fmt.Printf("%d %s by %s; ", b.ID, b.Title, b.Author)
		}
		fmt.Println()
		return
	}
	fmt.Println("No books???")
}

func genRandID() int {
	return rand.Int()
}

func genTimebasedId() int {
	return int(time.Now().UnixNano())
}

func main() {
	lib := storage.NewLibraryMap(genRandID)

	lib.AddBook("notebook", "")
	lib.AddBook("Urban Dynamics", "Jay W. Forrester")
	lib.AddBook("Hobbit, or There and Back Again", "J.R.R. Tolkien")
	printResponseSingle(lib.GetBookByTitle("Urban Dynamics"))

	lib.(*storage.StorageMap).GenerateID = genTimebasedId

	lib.AddBook("Notebook", "Some MIPT student")
	printResponseSingle(lib.GetBookByID(42))
	printResponseMulti(lib.Search("NoTeBOOk"))

	lib2 := storage.NewLibrarySlice(genRandID)
	lib2.AddBook("NotEbooK", "sus amogus")
	printResponseSingle(lib2.GetBookByTitle("Notebook"))

	library.Transfer(lib, lib2)
	lib.AddBook("Goblin Slayer", "") //Не будет в lib2
	printResponseMulti(lib2.Search("Notebook"))
	printResponseMulti(lib2.GetAll(), true)
}
