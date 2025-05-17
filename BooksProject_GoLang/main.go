package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	loadBooks()

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getAllBooks(w, r)
		case "POST":
			createBook(w, r)
		default:
			http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getBookbyId(w, r)
		case "PUT":
			updateBook(w, r)
		case "DELETE":
			deleteBook(w, r)
		default:
			http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/books/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
			return
		}
		searchBooks(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" //default pprt
	}

	fmt.Println("Server starting on PORT " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
