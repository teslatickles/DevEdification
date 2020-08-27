package controllers

import (
	"fmt"
	"github.com/DevEdification/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createVizNugInput struct {
	Title   string `json:"title" binding:"required"`
	Tech    string `json:"tech" binding:"required"`
	Company string `json:"company" binding:"required"`
	Author  string `json:"author" binding:"required"`
	GCSC    bool   `json:"gcsc"`
	URL     string `json:"url" binding:"required"`
}

type updateVizNugInput struct {
	Title   string `json:"title"`
	Tech    string `json:"tech"`
	Company string `json:"company"`
	Author  string `json:"author"`
	GCSC    bool   `json:"gcsc"`
	URL     string `json:"url"`
}

// FindVizNugs list all viznugs currently in db
func FindVizNugs(c *gin.Context) {
	var viznugs []models.VizNug
	models.DB.Find(&viznugs)

	c.JSON(http.StatusOK, gin.H{"data": viznugs})
}

// FindVizNug find specific viznug based on id
func FindVizNug(c *gin.Context) {
	var viznug models.VizNug

	if err := models.DB.Where("id = ?", c.Param("id")).First(&viznug).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": viznug})
}

// CreateVizNug add new viznug entry to viznug table
func CreateVizNug(c *gin.Context) {
	var input createVizNugInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//viznugSample := models.VizNug{
	//	Tech:   "Fortran",
	//	Title:  "Long ago...",
	//	Author: "Kanye Westpointe",
	//	GCSC:   false,
	//	URL:    "https://britisheyesonly.uk",
	//}

	viznug := models.VizNug{
		Tech:   input.Tech,
		Title:  input.Title,
		Author: input.Author,
		GCSC:   input.GCSC,
		URL:    input.URL,
	}

	models.DB.Create(&viznug)

	c.JSON(http.StatusOK, gin.H{"data": viznug})
}

// UpdateVizNug update specific viznug based on id
func UpdateVizNug(c *gin.Context) {
	var update updateVizNugInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var viznug models.VizNug
	if err := models.DB.Where("id = ?", c.Param("id")).First(&viznug).Update(&update).Error; err != nil {
		fmt.Print(c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DeleteVizNug delete specific viznug based on id
func DeleteVizNug(c *gin.Context) {
	var viznug models.VizNug
	if err := models.DB.Where("id = ?", c.Param("id")).First(&viznug).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&viznug)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
