package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// search function
func searchBooks(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		http.Error(w, "Search keyword is required", http.StatusBadRequest)
		return
	}

	keyword = strings.ToLower(keyword)

	goRnum := 4
	if len(books) < goRnum {
		goRnum = len(books)
	}

	if goRnum == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Book{})
		return
	}

	resultsChannel := make(chan []Book, goRnum)

	chunkSize := (len(books) + goRnum - 1) / goRnum

	for i := 0; i < goRnum; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if end > len(books) {
			end = len(books)
		}

		//if no books to process the skip
		if start >= end {
			resultsChannel <- []Book{} //sends an empty result
			continue
		}

		//!launching go routines for each chunk
		go func(chunk []Book) {
			var localResults []Book
			for _, book := range chunk {
				if strings.Contains(strings.ToLower(book.Title), keyword) ||
					strings.Contains(strings.ToLower(book.Description), keyword) {
					localResults = append(localResults, book)
				}
			}
			resultsChannel <- localResults
		}(books[start:end])
	}

	var results []Book
	for i := 0; i < goRnum; i++ {
		chunkResults := <-resultsChannel
		results = append(results, chunkResults...)
	}

	close(resultsChannel)

	w.Header().Set("Content-type", "application/json")

	if len(results) == 0 {
		http.Error(w, "Book not found", http.StatusBadRequest)
		return
	} else {
		json.NewEncoder(w).Encode(results)
	}

}
