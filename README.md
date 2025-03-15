# Book CRUD API in Go

A simple CRUD API for managing books, built using Go and a PostgreSQL database.

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
   git clone https://github.com/baonq888/go-bookstore.git
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
| POST   | `/book`       | Create a new book   |
| GET    | `/book/{id}`  | Retrieve a book     |
| PUT    | `/book/{id}`  | Update book details |
| DELETE | `/book/{id}`  | Delete a book       |

## Example Usage

### Create a Book
```sh
curl -X POST http://localhost:8080/book \
     -H "Content-Type: application/json" \
     -d '{"name": "Go Programming", "author": "John Doe", "publication": "Tech Books"}'
```

### Get a Book
```sh
curl -X GET http://localhost:8080/book/1
```

### Update a Book
```sh
curl -X PUT http://localhost:8080/book/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Advanced Go", "author": "Jane Doe", "publication": "Expert Publishing"}'
```

### Delete a Book
```sh
curl -X DELETE http://localhost:8080/book/1
```

## Running the Application

```sh
go run main.go
```

## Contributing
Feel free to fork this repository and submit pull requests.

## License
This project is licensed under the MIT License.

