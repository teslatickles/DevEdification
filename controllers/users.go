package controllers

import (
	"fmt"
	"github.com/DevEdification/v2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type createUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Role     string `json:"Role" binding:"required"`
}

type updateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"Email"`
	Role     string `json:"Role"`
}

// CreateUser creates a new user based on body
func CreateUser(c *gin.Context) {
	var input createUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Email: 	  input.Email,
		Role:	  input.Role,
	}
	models.DB.FirstOrCreate(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser updates user based on id
func UpdateUser(c *gin.Context) {
	var update updateUserInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Update(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// FindUser finds user based on id
func FindUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes user specified by id
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Login generate jwt for authorized user
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	fmt.Println(user.Role)
	if user.Role != "member" {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := createToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	c.JSON(http.StatusOK, token)
}

// createToken create then return valid jwt for authenticating user
func createToken(id uint) (string, error) {
	var err error
	// Create access token
	os.Setenv("ACCESS_SECRET", "boots")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
