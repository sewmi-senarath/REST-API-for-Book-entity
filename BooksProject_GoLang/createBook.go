package main

import (
	"encoding/json"
	"net/http"
)


// create books function
func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//!validations
	if newBook.BookId == "" {
		http.Error(w, "Book Id is required", http.StatusBadRequest)
		return
	}

	if newBook.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if newBook.Pages <= 0 {
		http.Error(w, "Invalid Page numbers", http.StatusBadRequest)
		return
	}

	if newBook.Price <= 0 {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	if newBook.Quantity <= 0 {
		http.Error(w, "Invalid Quantity", http.StatusBadRequest)
		return
	}

	//!check the ISBN validation
	if newBook.ISBN != "" && !validateISBN(newBook.ISBN) {
		http.Error(w, "ISBN value incorrect", http.StatusBadRequest)
		return
	}

	//!check for duplicate books
	for _, book := range books {
		if book.BookId == newBook.BookId {
			http.Error(w, "Book IDs cannot duplicate", http.StatusConflict)
			return
		}
	}

	//!check for duplicate isbn
	for _, book := range books {
		if book.ISBN == newBook.ISBN {
			http.Error(w, "ISBN cannot duplicate", http.StatusConflict)
			return
		}
	}

	books = append(books, newBook)
	saveBooks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}