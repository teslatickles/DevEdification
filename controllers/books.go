package controllers

import (
	"github.com/DevEdification/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBookInput
type CreateBookInput struct {
	Title   string `json:"title" binding:"required"`
	Release string `json:"release" binding:"required"`
	Author  string `json:"author" binding:"required"`
	URL     string `json:"url" binding:"required"`
}

// UpdateBookInput
type UpdateBookInput struct {
	Title   string `json:"title"`
	Release string `json:"release"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}

// FindBooks
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:   input.Title,
		Release: input.Release,
		Author:  input.Author,
		URL:     input.URL,
	}
	models.DB.FirstOrCreate(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook
func UpdateBook(c *gin.Context) {
	var update UpdateBookInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Update(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DeleteBook
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
