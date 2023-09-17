# Go Fiber PostgreSQL API
This is a simple RESTful API built using the Go Fiber framework, PostgreSQL database, and GORM ORM for handling book-related operations. The API allows you to create, retrieve, update, and delete book records in a PostgreSQL database. It also provides pagination for fetching a list of books.
# Prerequisites
Before you can run this API, make sure you have the following installed:
1. Go (v1.16 or higher)
2. PostgreSQL
3. GORM (Go Object Relational Mapping)
4. Go Fiber (Fiber v2)
5. github.com/joho/godotenv for handling environment variables
# Getting Started
1. Clone the repository to your local machine:
git clone https://github.com/bakhtybayevn/postgreSql.git
2. Create a .env file in the project root and set the following environment variables:
DB_HOST=your_postgresql_host
DB_PORT=your_postgresql_port
DB_USER=your_postgresql_user
DB_PASS=your_postgresql_password
DB_NAME=your_postgresql_database
DB_SSLMODE=disable
Modify the values according to your PostgreSQL database configuration.
3. Install the required Go packages:
go mod tidy
4. Migrate the database to create the necessary tables:
go run main.go migrate
# Usage
1. To start the API, run the following command::
go run main.go
The API will start and be accessible at http://localhost:8080.
# Endpoints
1. Create Book: POST /api/create_books
   Request Body (JSON):
   {
  "author": "Author Name",
  "title": "Book Title",
  "publisher": "Book Publisher"
   }
2. Delete Book: DELETE /api/delete_book/:id
Replace :id with the ID of the book you want to delete.
3. Get Book by ID: GET /api/get_books/:id
Replace :id with the ID of the book you want to retrieve.
4. Get Books (Paginated): GET /api/books?page=<page_number>&limit=<items_per_page>
Use query parameters page and limit for pagination.
5. Get All Books: GET /api/books_all
# Sample Requests
Create Book:
curl -X POST http://localhost:8080/api/create_books -d '{
  "author": "J.K. Rowling",
  "title": "Harry Potter and the Sorcerer's Stone",
  "publisher": "Scholastic"
}'

Delete Book:
curl -X DELETE http://localhost:8080/api/delete_book/1

Get Book by ID:
curl http://localhost:8080/api/get_books/1

Get Books (Paginated):
curl http://localhost:8080/api/books?page=1&limit=10

Get All Books:
curl http://localhost:8080/api/books_all
# Acknowledgments
1. Go Fiber - A web framework for Go.
2. GORM - The Go Object Relational Mapping library.
3. PostgreSQL - A powerful, open-source relational database system.
