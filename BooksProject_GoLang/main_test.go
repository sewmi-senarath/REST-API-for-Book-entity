package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBookById(t *testing.T) {
	//!this is the book we test
	books = []Book{
		{
			BookId:          "1",
			AuthorId:        "e0d91f68-a183-477d",
			PublisherId:     "2f7b19e9-b268-4440-a15b",
			Title:           "Alice in Wonderful",
			PublicationDate: "1925-04-10",
			ISBN:            "9780743273578",
			Pages:           10,
			Genre:           "Novel",
			Description:     "Mystery Book",
			Price:           200.00,
			Quantity:        10,
		},
	}

	//!Test 01
	req1, _ := http.NewRequest("GET", "/books/1", nil)
    rr1 := httptest.NewRecorder()
    getBookbyId(rr1, req1)

    if rr1.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", rr1.Code)
    }
    var book Book
    json.Unmarshal(rr1.Body.Bytes(), &book)
    if book.BookId != "1" {
        t.Errorf("Expected book ID '1', got '%s'", book.BookId)
    }

	//!test 02
	// Test 2: Non-existing book
    req2, _ := http.NewRequest("GET", "/books/2", nil)
    rr2 := httptest.NewRecorder()
    getBookbyId(rr2, req2)

    if rr2.Code != http.StatusNotFound {
        t.Errorf("Expected status 404, got %d", rr2.Code)
    }
    if rr2.Body.String() != "Book not found\n" {
        t.Errorf("Expected 'Book not found\\n', got '%s'", rr2.Body.String())
    }


}
