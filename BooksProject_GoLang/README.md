
**Books Project (GoLang)**

## Overview
    A simple RESTful API built to manage a collection of books, including CRUD operations, search functionality, Unit Test case for the 'getBookById' function

## Features
    1. Create a new book
        Validations for ISBN, BookId for uniqueness
        Validations for the page numbers, Price and Quantity
    2. Update the books preserving the same validations
    3. View all books with pagination
    4. View a single book by ID
    5. Delete a book by ID
    6. Search book by keyword which is case-insensitive
    7. Unit test for the 'getBookById' function

## Requirements
    - Go: Version 1.24 or higher 
    - Docker Desktop: Required for containerization and Minikube’s Docker driver. 
    - Minikube: For local Kubernetes deployment. Install via Chocolatey: `choco install minikube -y`.
    - kubectl: Kubernetes CLI. Install via Chocolatey: `choco install kubernetes-cli -y`.
    - PowerShell: Windows PowerShell for running commands

## Installation
    1. Clone the Project
    2. Install Go
    3. Install Docker Desktop
    4. Install Minikube and kubectl (for Kubernetes)

## Usage
    1. Using Go:
        go run (Get-ChildItem *.go | Where-Object { $_.Name -notlike "*_test.go" })
        http://localhost:8081

    2. Using Docker:
        docker build -t books-api .
        docker run -p 8081:8081 books-api
        http://localhost:8081

    3. Using Kubernetes:
        - Prerequisites
            Docker Desktop: Installed and running (for the Docker driver).
            Minikube: Install via Chocolatey: choco install minikube -y.
            kubectl: Install via Chocolatey: choco install kubernetes-cli -y.
            PowerShell: Run commands in an Administrator PowerShell session.
        - Steps to Deploy
            Start Minikube:
                minikube start --driver=docker
            Build the Docker Image in Minikube:
                & minikube -p minikube docker-env | Invoke-Expression
            Build the image:
                docker build -t books-api .
            Deploy the Application:
                kubectl apply -f deployment.yaml
                kubectl apply -f service.yaml
            Access the Application:
                minikube service books-api-service --url
            Test in a new PowerShell:
                Invoke-RestMethod -Uri "<service url>" -Method Get

## Running tests
    go test -v


## API endpoints
    GET	    /books	        Get all books (paginated)	
    POST	/books	        Create a new book	
    GET	    /books/{id}	    Get a book by ID	
    PUT	    /books/{id}	    Update a book by ID	
    DELETE	/books/{id}	    Delete a book by ID	
