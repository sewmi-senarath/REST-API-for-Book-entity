package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Book struct{
	BookId string `json:"bookId"`
	AuthorId string `json:"authorId"`
	PublisherId string `json:"publisherId"`
	Title string `json:"title"`
	PublicationDate string `json:"publicationDate"`
	ISBN string `json:"isbn"`
	Pages int `json:"pages"`
	Genre string `json:"genre"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

var books []Book
const bookFilePath = "books.json"

func loadBooks(){
	data, error := ioutil.ReadFile(bookFilePath)
	if error != nil{
		books = []Book{}
		return
	}
	json.Unmarshal(data, &books)
}

func saveBooks(){
	data, _ := json.MarshalIndent(books, "", " ")
	ioutil.WriteFile(bookFilePath, data, 0644)
}

// !ISBN validations checking
// ! 13 digits  , seperated by hyphen and space
func validateISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn = strings.ReplaceAll(isbn, " ", "")

	if len(isbn) != 13 {
		return false
	}

	for _, char := range isbn {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
