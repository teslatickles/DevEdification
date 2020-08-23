package controllers

import (
	"github.com/DevEdification/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	models.DB.FirstOrCreate(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
