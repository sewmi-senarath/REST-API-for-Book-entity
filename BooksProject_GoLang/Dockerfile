FROM golang:1.24

# Set working directory inside the container
WORKDIR /app

# Copy go.mod
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy all Go source files
COPY *.go ./

COPY books.json ./

# Build the Go application
RUN go build -o books-api .

# Expose port 8081
EXPOSE 8081

# Run the application
CMD ["./books-api"]