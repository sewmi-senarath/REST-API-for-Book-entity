package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// delete books function
func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")

	if id == "" {
		http.Error(w, "bookId is required in URL", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.BookId == id {
			books = append(books[:i], books[i+1:]...)
			saveBooks()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Book deleted successfully",
			})
			return
		}
	}
	http.Error(w, "Book not found to delete", http.StatusNotFound)
}