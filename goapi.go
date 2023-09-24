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

func updateBookById(w http.ResponseWriter, req *http.Request) {
	var book Book
	err := json.NewDecoder(req.Body).Decode(&book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books[book.ID] = book
	fmt.Fprintf(w, "%v", book)
}

func deleteBookById(w http.ResponseWriter, req *http.Request) {
	var inputData map[string]string
	err := json.NewDecoder(req.Body).Decode(&inputData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, ok := inputData["id"]

	if !ok {
		http.Error(w, "Missing 'id' field in JSON data", http.StatusBadRequest)
		return
	}

	delete(books, id)
	fmt.Fprintf(w, "book with the id:%s is deleted", id)
}

func main() {
	http.HandleFunc("/books/all", getAllBooks)
	http.HandleFunc("/books/add", addBook)
	http.HandleFunc("/books/update", updateBookById)
	http.HandleFunc("/books/delete", deleteBookById)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
