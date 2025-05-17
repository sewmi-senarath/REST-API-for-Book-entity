package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"strconv"
)

// get all books function
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	//!getting query parameters
	limitString := r.URL.Query().Get("limit")
	offsetString := r.URL.Query().Get("offset")

	//!default values
	limit := 10
	offset := 0

	if limitString != "" {
        if l, err := strconv.Atoi(limitString); err == nil && l > 0 {
            limit = l
        } 
    }

    if offsetString != "" {
        if o, err := strconv.Atoi(offsetString); err == nil && o >= 0 {
            offset = o
        } 
    }

	//!if os > total, it gives an emty array
    if offset >= len(books) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode([]Book{})
        return
    }

    end := offset + limit
    if end > len(books) {
        end = len(books)
    }

	paginatedBooks := books[offset:end]

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(paginatedBooks)
}

// get book by id function
func getBookbyId(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	for _, book := range books {
		if book.BookId == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)

}