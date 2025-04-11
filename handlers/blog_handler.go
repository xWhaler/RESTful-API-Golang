package handlers

import (
	"database/sql"
	"net/http"
	"go-rest/models"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	DB *sql.DB
}

func (h *BlogHandler) GetPosts(c *gin.Context) {

	b, err := models.GetPosts(h.DB)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, b)
}

func (h *BlogHandler) CreatePost(c *gin.Context) {
	var b models.BlogPost
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreatePost(h.DB, b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, b)
}
