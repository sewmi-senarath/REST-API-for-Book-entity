REST API for Book entity
This project implements a REST API for managing books, built with GoLang, as per the requirements of the Software Engineering Intern Technical Assignment. The application uses a text file (JSON format) as the data persistence layer and includes CRUD operations, a keyword search endpoint, and performance optimization using Go's concurrency primitives. Bonus tasks, such as Docker containerization and unit tests, are also included.
Prerequisites
To run this application, ensure you have the following installed:

Go: Version 1.21 or later (https://go.dev/doc/install)
Docker: (Optional, for containerization) (https://docs.docker.com/get-docker/)
A text editor or IDE (e.g., VSCode)


Setup and Running the Application
Local Setup

Clone or Unzip the Project:Extract the project zip file to your desired directory.

Navigate to the Project Directory:
cd path/to/project


Install Dependencies:The project uses standard Go libraries and the gorilla/mux router. Install it using:
go mod tidy
go get github.com/gorilla/mux


Run the Application:Start the server with:
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

Example Requests

Get All Books:
curl http://localhost:8080/books


Create a Book:
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866", "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70", "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607", "title": "The Great Gatsby", "publicationDate": "1925-04-10", "isbn": "9780743273565", "pages": 180, "genre": "Novel", "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.", "price": 15.99, "quantity": 5}'


Search Books:
curl http://localhost:8080/books/search?q=gatsby



Data Storage
The application uses a books.json file as the persistence layer. The file is automatically created if it doesn't exist when the application starts. Ensure the application has write permissions in the directory containing books.json.
Search Optimization
The search endpoint (GET /books/search) is optimized using:

Goroutines: Books are split into subsets, and each subset is searched concurrently.
Channels: Results from each goroutine are collected via a channel and merged into a single response.

Bonus Features Implemented

Docker Containerization:

A Dockerfile is included to containerize the application.
Instructions for building and running the Docker image are provided above.


Unit Tests:

Unit tests for the GET /books/{id} endpoint are implemented in tests/book_test.go.
Run tests with:go test ./tests




Pagination:

The GET /books endpoint supports pagination using query parameters limit and offset.
Example: GET /books?limit=10&offset=20


Kubernetes Deployment:

Kubernetes manifest files are not included in this submission due to time constraints, but the application is compatible with Kubernetes. To deploy, create a deployment and service manifest, and use Minikube or Kind as specified.



Running Tests
To execute the unit tests:
go test ./tests -v


Troubleshooting

Port Conflict: If port 8080 is in use, change the port in main.go or stop the conflicting process.
File Permissions: Ensure the application has write access to create or modify books.json.
Dependency Issues: Run go mod tidy to resolve missing dependencies.
