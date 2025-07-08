package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-rest/handlers"
	"log"
)

func main() {
	// Initialize database connection
	db, err := sql.Open("mysql", "[root]:[password]@tcp(127.0.0.1:3306)/[db_name]")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t := db.Ping()
	if t != nil {
		log.Fatal(t)
	}

	// Initialize blog handler
	blogHandler := &handlers.BlogHandler{DB: db}

	// Initialize router
	router := gin.Default()

	// Load HTML templates from the /templates directory
	router.LoadHTMLGlob("templates/*")

	// Setup routes
	setupRoutes(router, blogHandler, db)

	// Start the server
	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, blogHandler *handlers.BlogHandler, db *sql.DB) {
	// Route to show all posts
	router.GET("/", blogHandler.GetPosts)

	// Route to show individual post by ID
	router.GET("/post/:id", blogHandler.GetPostByID)
}
