package main

import (
	"github.com/DevEdification/v2/controllers"
	//_ "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	//_ "github.com/MartinHeinz/go-project-blueprint/pkg"

	//"github.com/swaggo/files"
	//_ "github.com/MartinHeinz/go-project-blueprint"
	_ "github.com/DevEdification/v2/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

// @title Amozone API
// @version 2.0
// @description Swagger page for Amozone Golang API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email hunterhartline87@gmail.com

// @license.name MIT
// @license.url "/LICENSE"

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func initRoutes() {
	// Init gin engine
	r := initGin()

	// Swagger API route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Book routes
	r.GET("/books/:id", controllers.FindBook)
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Website routes
	r.GET("/websites/:id", controllers.FindWebsite)
	r.GET("/websites", controllers.FindWebsites)
	r.POST("/websites", controllers.CreateWebsite)
	r.PATCH("/websites/:id", controllers.UpdateWebsite)
	r.DELETE("/websites/:id", controllers.DeleteWebsite)

	// VizNug routes
	r.GET("/viznugs/:id", controllers.FindVizNug)
	r.GET("/viznugs", controllers.FindVizNugs)
	r.POST("/viznugs", controllers.CreateVizNug)
	r.PATCH("/viznugs/:id", controllers.UpdateVizNug)
	r.DELETE("/viznugs/:id", controllers.DeleteVizNug)

	// User routes
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/login", controllers.Login)

	// attach router to server - handle errors
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}

}

// initGin initialize gin engine with default configuration
// set landing page and return router variable r
func initGin() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to Vizient's Software Wizard Manual"})
	})

	return r
}
