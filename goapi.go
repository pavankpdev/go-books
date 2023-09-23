package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = make(map[string]Book)

func addBook(w http.ResponseWriter, req *http.Request) {
	var book Book
	err := json.NewDecoder(req.Body).Decode(&book)
	fmt.Println(book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books[book.ID] = book

	fmt.Fprintf(w, "%v", book)
}

func getAllBooks(w http.ResponseWriter, req *http.Request) {
	var allBook []Book

	for _, value := range books {
		allBook = append(allBook, value)
	}

	fmt.Fprintf(w, "%v", allBook)
}

func main() {
	http.HandleFunc("/books/", getAllBooks)
	http.HandleFunc("/books/add", addBook)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
