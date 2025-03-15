# Book CRUD API in Go

A simple CRUD (Create, Read, Update, Delete) API for managing books, built using Go and a PostgreSQL database.

## Features
- Create a new book
- Retrieve a book by ID
- Update book details
- Delete a book

## Prerequisites
- Go installed (1.18+ recommended)
- PostgreSQL database
- `github.com/lib/pq` for PostgreSQL driver (if using PostgreSQL)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/book-crud-api.git
   cd book-crud-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up your database and update the connection string in the code.

## API Endpoints

| Method | Endpoint        | Description          |
|--------|----------------|----------------------|
| POST   | `/books`       | Create a new book   |
| GET    | `/books/{id}`  | Retrieve a book     |
| PUT    | `/books/{id}`  | Update book details |
| DELETE | `/books/{id}`  | Delete a book       |

## Example Usage

### Create a Book
```sh
curl -X POST http://localhost:8080/books \
     -H "Content-Type: application/json" \
     -d '{"name": "Go Programming", "author": "John Doe", "publication": "Tech Books"}'
```

### Get a Book
```sh
curl -X GET http://localhost:8080/books/1
```

### Update a Book
```sh
curl -X PUT http://localhost:8080/books/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Advanced Go", "author": "Jane Doe", "publication": "Expert Publishing"}'
```

### Delete a Book
```sh
curl -X DELETE http://localhost:8080/books/1
```

## Running the Application

```sh
go run main.go
```

## Contributing
Feel free to fork this repository and submit pull requests.

## License
This project is licensed under the MIT License.

