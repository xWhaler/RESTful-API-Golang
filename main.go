package main

import (
	"database/sql"
	"go-rest/handlers"
	"go-rest/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "camelbyte:Sophia@tcp(127.0.0.1:3306)/goblog")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	blogHandler := &handlers.BlogHandler{DB: db}

	router := gin.Default()

	router.GET("/posts", blogHandler.GetPosts)

	router.POST("/templates/create", func(c *gin.Context) {
		var newPost models.BlogPost

		if err := c.Bind(&newPost); err != nil {
			c.String(http.StatusBadRequest, "Bad Request: %v", err)
			return
		}

		err = models.CreatePost(db, newPost) // Pass newPost to CreatePost
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating the blog post: %v", err)
			return
		}

		c.String(http.StatusOK, "Blog post created successfully")
	})

	router.LoadHTMLGlob("templates/*")

	router.Run(":8080") // Ensure this is outside all handlers
}
