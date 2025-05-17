Book Entity API

This project implements a REST API for managing books, built with GoLang, as per the requirements. The application uses a text file (JSON format) as the data persistence layer and includes CRUD operations, a keyword search endpoint, and performance optimization using Go's concurrency primitives. Bonus tasks, such as Docker containerization and unit tests, are also included.

Prerequisites

To run this application, ensure you have the following installed:





Go: Version 1.21 or later (https://go.dev/doc/install)



Docker: (Optional, for containerization) (https://docs.docker.com/get-docker/)



A text editor or IDE (e.g., VSCode)
Setup and Running the Application

Local Setup





Clone or Unzip the Project: Extract the project zip file to your desired directory.



Navigate to the Project Directory:

cd path/to/project



Install Dependencies: The project uses standard Go libraries and the gorilla/mux router. Install it using:

go mod tidy
go get github.com/gorilla/mux



Run the Application: Start the server with:

go run main.go

The API will be available at http://localhost:8080.

Running with Docker





Build the Docker Image:

docker build -t book-api .



Run the Docker Container:

docker run -p 8080:8080 book-api

The API will be accessible at http://localhost:8080.

API Endpoints

The API supports the following endpoints:





GET /books: Retrieve a list of all books (supports pagination with ?limit=<int>&offset=<int>).



POST /books: Create a new book (expects JSON payload).



GET /books/{id}: Retrieve a book by its ID.



PUT /books/{id}: Update a book by its ID (expects JSON payload).



DELETE /books/{id}: Delete a book by its ID.



GET /books/search?q=: Search books by keyword in title or description (case-insensitive).
