# Go-Rest Blog API

A simple RESTful API built with Go and Gin that allows you to manage blog posts. It supports creating new posts and retrieving all posts.


#### <b> Working on adding Authentication, TLS, Users. PULL REQUEST</b>

## Features

- **GET /blogs/**: Retrieve all blog posts.
- **POST /blogs/**: Create a new blog post.

## Technologies Used

- **Go**: Programming language used to build the API.
- **Gin**: Web framework for building the REST API.
- **SQL (MariaDB/MySQL)**: Database for storing blog posts.
- **Go Modules**: Dependency management for Go projects.

## Project Structure

```plaintext
/go-rest
  ├── main.go           # Main entry point for the application
  ├── handlers/         # Contains request handlers for blog operations
  │   └── blog_handler.go
  ├── models/           # Database models and interactions
  │   └── blog.go
  ├── routes/           # Routes for API endpoints
  │   └── blog_routes.go
  ├── static/           # Static files (CSS, JS, images)
  │   └── css/
  │   └── javascript/
  ├── templates/        # HTML templates (if rendering HTML)
  ├── migrations/       # Database migrations (optional)
  └── api/              # API-specific files (optional)
      └── blog_api.go

Setup Instructions
Prerequisites

    Go (version 1.16 or higher)

    MariaDB or MySQL (or any compatible SQL database)

    A Git client (if cloning the repository)

Step 1: Clone the Repository

Clone the repository to your local machine:

git clone https://github.com/yourusername/go-rest.git
cd go-rest

Step 2: Install Dependencies

Make sure Go modules are enabled and install the dependencies:

go mod tidy

Step 3: Set Up the Database

Create a database (e.g., mydb) in MariaDB or MySQL and set up the blog_posts table:

CREATE DATABASE mydb;

USE mydb;

CREATE TABLE blog_posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(255),
    content TEXT
);

Step 4: Configure Database Connection

Update your database connection settings in the main.go or the relevant config file (if applicable) to point to your MariaDB/MySQL server.

Example (in main.go):

// Update the database connection string
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")

Step 5: Run the Application

To build and run the application:

go build
./go-rest

The server will start and listen on localhost:8080. You can access the API:

    GET /blogs/: Fetch all blog posts.

    POST /blogs/: Create a new blog post. (Example JSON body for POST request)

{
  "title": "My First Blog Post",
  "author": "John Doe",
  "content": "This is the content of the blog post."
}

Step 6: Testing the API

You can use tools like Postman or curl to test the API:

Get all posts:

curl http://localhost:8080/blogs/

Create a new post:

curl -X POST http://localhost:8080/blogs/ \
  -H "Content-Type: application/json" \
  -d '{"title": "New Blog Post", "author": "Jane Doe", "content": "This is a new post."}'

License

This project is licensed under the MIT License - see the LICENSE file for details.


### Things to customize:
- Replace `yourusername` in the GitHub clone URL.
- Add any specific setup or configuration details.
- Mention any additional dependencies or setup requirements if needed.

Let me know if you'd like further changes or additions!
