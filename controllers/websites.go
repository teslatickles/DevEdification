package controllers

import (
	"github.com/DevEdification/v2/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
)

// createWebsiteInput input struct used for creating
// new website entry in db
type createWebsiteInput struct {
	Title   string `json:"title" binding:"required"`
	Tech    string `json:"tech" binding:"required"`
	Company string `json:"company" binding:"required"`
	Author  string `json:"author" binding:"required"`
	URL     string `json:"url" binding:"required"`
}

// updateWebsiteInput input struct used for
// updating existing website entry with new values
type updateWebsiteInput struct {
	Title   string `json:"title"`
	Tech    string `json:"tech"`
	Company string `json:"company"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}

// FindWebsites godoc
// @Summary Find all websites
// @Description retrieve all website entries
// @ID get-list
// @Accept json
// @Produce json
// @Success 200 {object} models.Website
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /websites/ [get]
// FindWebsites list all website entries contained in database
func FindWebsites(c *gin.Context) {
	var websites []models.Website
	models.DB.Find(&websites)

	c.JSON(http.StatusOK, gin.H{"data": websites})
}

// FindWebsite godoc
// @Summary Find a website by ID
// @Description find a website entry by ID
// @ID get-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.Website
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /websites/{id} [get]
// FindWebsite find website entry based on id
func FindWebsite(c *gin.Context) {
	var website models.Website

	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": website})
}

// CreateWebsite godoc
// @Summary Create website entry
// @Description create a new website entry
// @ID create-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.Website
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /websites/{id} [post]
// CreateWebsite create new website
func CreateWebsite(c *gin.Context) {
	var input createWebsiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	website := models.Website{
		Title:   input.Title,
		Tech:    input.Tech,
		Company: input.Company,
		Author:  input.Author,
		URL:     input.URL,
	}
	models.DB.FirstOrCreate(&website)

	c.JSON(http.StatusOK, gin.H{"data": website})
}

// UpdateWebsite godoc
// @Summary Update website entry
// @Description update an existing website entry
// @ID update-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.Website
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /websites/{id} [patch]
// UpdateWebsite update website based on id
func UpdateWebsite(c *gin.Context) {
	var update updateWebsiteInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var website models.Website
	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Update(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &update})
}

// DeleteWebsite godoc
// @Summary Delete website entry
// @Description delete an existing website entry
// @ID delete-entry
// @Accept json
// @Produce json
// @Success 200 {object} models.Website
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /websites/{id} [delete]
// DeleteWebsite delete website entry based on id
func DeleteWebsite(c *gin.Context) {
	var website models.Website
	if err := models.DB.Where("id = ?", c.Param("id")).First(&website).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&website)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
