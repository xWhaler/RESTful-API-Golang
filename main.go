package main 

import (
    "database/sql"
    "log"
	"go-rest/handlers"
    
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)


// This is a simple REST API for a blog application using Go and Gin framework.

func main() {

	blogdb, err := sql.Open("mysql", "camelbyte:Sophia@tcp(127.0.0.1:3306)/goblog")
	if err != nil {
		log.Fatal(err)
	}
	defer blogdb.Close()

	blogHandler := &handlers.BlogHandler{DB: blogdb}

	router := gin.Default()

	router.GET("/posts", blogHandler.GetPosts)

	router.POST("/", blogHandler.CreatePost)
	

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ponging back from the ping"})
		})

	router.Run(":8080")
}
