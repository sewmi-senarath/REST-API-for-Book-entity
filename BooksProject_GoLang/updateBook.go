package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// update books function
func updateBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if id == "" {
		http.Error(w, "Book id is required for updates", http.StatusBadRequest)
		return
	}

	var bookUpdate *Book
	var index int
	for i, book := range books {
		if book.BookId == id {
			bookUpdate = &books[i]
			index = i
			break
		}
	}
	if bookUpdate == nil {
		http.Error(w, "Book not found to update", http.StatusNotFound)
		return
	}

	var updates map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Ïnvalid request type", http.StatusBadRequest)
		return
	}

	if bookId, ok := updates["bookId"]; ok {
		if v, ok := bookId.(string); ok && v != id {
			http.Error(w, "Book ID cannot update", http.StatusBadRequest)
			return
		}
	}

	//!for the partial updates of individual fields
	for key, value := range updates {
		switch key {
		case "bookId":
			if v, ok := value.(string); ok {
				if v != id {
					for _, book := range books {
						if book.BookId == v && book.BookId != id {
							http.Error(w, "Book ID already exists", http.StatusConflict)
							return
						}
					}
					bookUpdate.BookId = v
				}
			}
		case "authorId":
			if v, ok := value.(string); ok {
				bookUpdate.AuthorId = v
			}
		case "publisherId":
			if v, ok := value.(string); ok {
				bookUpdate.PublisherId = v
			}
		case "title":
			if v, ok := value.(string); ok {
				bookUpdate.Title = v
			}
		case "publicationDate":
			if v, ok := value.(string); ok {
				bookUpdate.PublicationDate = v
			}
		case "isbn":
			if v, ok := value.(string); ok {
				if v != "" && !validateISBN(v) {
					http.Error(w, "ISBN value incorrect", http.StatusBadRequest)
					return
				}
				if v != "" && v != bookUpdate.ISBN {
					for _, book := range books {
						if book.ISBN == v && book.BookId != id {
							http.Error(w, "ISBN already exists", http.StatusConflict)
							return
						}
					}
				}
				bookUpdate.ISBN = v
			}
		case "pages":
			if v, ok := value.(float64); ok {
				if v <= 0 {
					http.Error(w, "pages cannot be negative or 0", http.StatusBadRequest)
					return
				}
				bookUpdate.Pages = int(v)
			}
		case "genre":
			if v, ok := value.(string); ok {
				bookUpdate.Genre = v
			}
		case "description":
			if v, ok := value.(string); ok {
				bookUpdate.Description = v
			}
		case "price":
			if v, ok := value.(float64); ok {
				if v <= 0 {
					http.Error(w, "prices cannot be negative or 0", http.StatusBadRequest)
					return
				}
				bookUpdate.Price = v
			}
		case "quantity":
			if v, ok := value.(float64); ok {
				if v <= 0 {
					http.Error(w, "quantity cannot be negative", http.StatusBadRequest)
					return
				}
				bookUpdate.Quantity = int(v)
			}

		}
	}

	books[index] = *bookUpdate
	saveBooks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookUpdate)
}
