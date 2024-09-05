# Go CRUD API for Bookstore Management System
This mini-project is a simple CRUD (Create, Read, Update, Delete) API built using Golang. <br> It allows users to manage a collection of books stored in a MySQL database, following a tutorial from freeCodeCamp.org.

## Features
__Create__: Add a new book to the collection in the database.<br>
__Read__: Retrieve all books or a specific book by ID.<br>
__Update__: Modify an existing book's details.<br>
__Delete__: Remove a book from the collection in the database.<br>

## Project Structure
__cmd/main__: Contains the main.go file, which is the entry point of the application and listens on the localhost for incoming API requests.<br>
__pkg/config__: Holds the configuration files, including .env and app.go, which is responsible for setting up the environment variables and establishing a connection to the database.<br>
__pkg/controllers__: Contains the controller logic, like book-controller.go, which handles the API requests and communicates with the services for performing CRUD operations on books.<br>
__pkg/models__: Defines the Book struct and its attributes, such as ID, Title, Author, and Publisher.<br>
__pkg/routes__: Handles the routing for the API requests and maps the endpoints to their corresponding controller methods in bookstore-routes.go.<br>
__pkg/utils__: Includes utility functions like input validation or helper methods that may be used across the project.<br>

## API Endpoints
POST /books - Add a new book to the collection.<br>
GET /books - Retrieve all books.<br>
GET /books/{id} - Retrieve a book by its ID.<br>
PUT /books/{id} - Update an existing book by its ID.<br>
DELETE /books/{id} - Delete a book by its ID.<br>

## Tutorial Reference
This project is based on the freeCodeCamp.org tutorial:⌨️ Golang With MYSQL Book Management System
