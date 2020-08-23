package controllers

import (
	"github.com/DevEdification/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateWebsiteInput
type CreateWebsiteInput struct {
	Title 	string	`json:"title" binding:"required"`
	Tech 	string 	`json:"tech" binding:"required"`
	Company string  `json:"company" binding:"required"`
	Author 	string 	`json:"author" binding:"required"`
	URL		string 	`json:"url" binding:"required"`
}

// UpdateWebsiteInput
type UpdateWebsiteInput struct {
	Title 	string	`json:"title"`
	Tech 	string  `json:"tech"`
	Company	string  `json:"company"`
	Author 	string 	`json:"author"`
	URL		string 	`json:"url"`
}

// FindWebsites
func FindWebsites(c *gin.Context) {
	var websites []models.Website
	models.DB.Find(&websites)

	c.JSON(http.StatusOK, gin.H{"data": websites})
}

// FindWebsite
func FindWebsite(c *gin.Context) {
	var website models.Website

	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": website})
}

// CreateWebsite
func CreateWebsite(c *gin.Context) {
	var input CreateWebsiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	website := models.Website{
		Title:  input.Title,
		Tech: input.Tech,
		Company: input.Company,
		Author: input.Author,
		URL:    input.URL,
	}
	models.DB.FirstOrCreate(&website)

	c.JSON(http.StatusOK, gin.H{"data": website})
}

// UpdateWebsite
func UpdateWebsite(c *gin.Context) {
	var update UpdateWebsiteInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var website models.Website
	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Update(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data":&update})
}

// DeleteWebsite
func DeleteWebsite(c *gin.Context) {
	var website models.Website
	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&website)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
