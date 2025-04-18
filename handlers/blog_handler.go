package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	DB *sql.DB
}

type BlogPost struct {
	ID        int
	Title     string
	Author    string
	Content   string
	CreatedOn string
	Tags      string
	Subject   string
}

func (h *BlogHandler) GetPosts(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, title, author, content, created_on, tags, subject FROM BlogPosts ORDER BY created_on DESC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading posts")
		return
	}
	defer rows.Close()

	var posts []BlogPost
	for rows.Next() {
		var post BlogPost
		err := rows.Scan(&post.ID, &post.Title, &post.Author, &post.Content, &post.CreatedOn, &post.Tags, &post.Subject)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing row")
			return
		}
		posts = append(posts, post)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Blog Posts",
		"Posts": posts,
	})
}

func (h *BlogHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post BlogPost
	err := h.DB.QueryRow("SELECT id, title, author, content, created_on, tags, subject FROM BlogPosts WHERE id = ?", id).
		Scan(&post.ID, &post.Title, &post.Author, &post.Content, &post.CreatedOn, &post.Tags, &post.Subject)
	if err != nil {
		c.String(http.StatusNotFound, "Post not found")
		return
	}

	c.HTML(http.StatusOK, "post.html", gin.H{
		"Post": post,
	})
}
