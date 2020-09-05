package controllers

import (
	"fmt"
	"github.com/DevEdification/v2/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
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

// FindVizNugs godoc
// @Summary Find all viznugs
// @Description retrieve all viznug entries
// @ID get-list
// @Accept json
// @Produce json
// @Success 200 {object} models.VizNug
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /viznugs/ [get]
// FindVizNugs list all viznugs currently in db
func FindVizNugs(c *gin.Context) {
	var viznugs []models.VizNug
	models.DB.Find(&viznugs)

	c.JSON(http.StatusOK, gin.H{"data": viznugs})
}

// FindVizNug godoc
// @Summary Find a viznug
// @Description retrieve a viznug based on ID
// @ID get-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.VizNug
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /viznugs/{id} [get]
// FindVizNug find specific viznug based on id
func FindVizNug(c *gin.Context) {
	var viznug models.VizNug

	if err := models.DB.Where("id = ?", c.Param("id")).First(&viznug).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": viznug})
}

// CreateVizNug godoc
// @Summary Create a viznug
// @Description create a viznug based on ID
// @ID create-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.VizNug
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /viznugs/{id} [post]
// CreateVizNug add new viznug entry to viznug table
func CreateVizNug(c *gin.Context) {
	var input createVizNugInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

// UpdateVizNug godoc
// @Summary Update a viznug
// @Description update a viznug based on ID
// @ID update-entry
// @Accept json
// @Produce json
// @Param id path int true "id of viznug entry to update"
// @Success 200 {object} models.VizNug
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /viznugs/{id} [patch]
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

// DeleteVizNug godoc
// @Summary Delete a viznug
// @Description delete a viznug based on ID
// @ID delete-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.VizNug
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /viznugs/{id} [delete]
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
