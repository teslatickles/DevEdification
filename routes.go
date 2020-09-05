package main

import (
	"github.com/DevEdification/v2/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	_ "github.com/DevEdification/v2/docs"
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

	//c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		// Book routes
		books := v1.Group("/books")
		{
			books.GET(":id", controllers.FindBook)
			books.GET("", controllers.FindBooks)
			books.POST("", controllers.CreateBook)
			books.PATCH(":id", controllers.UpdateBook)
			books.DELETE(":id", controllers.DeleteBook)
		}
		// Website routes
		websites := v1.Group("/websites")
		{
			websites.GET(":id", controllers.FindWebsite)
			websites.GET("", controllers.FindWebsites)
			websites.POST("", controllers.CreateWebsite)
			websites.PATCH(":id", controllers.UpdateWebsite)
			websites.DELETE(":id", controllers.DeleteWebsite)
		}
		// VizNug routes
		viznugs := v1.Group("/viznugs")
		{
			viznugs.GET(":id", controllers.FindVizNug)
			viznugs.GET("", controllers.FindVizNugs)
			viznugs.POST("", controllers.CreateVizNug)
			viznugs.PATCH(":id", controllers.UpdateVizNug)
			viznugs.DELETE(":id", controllers.DeleteVizNug)
		}
		// User routes
		users := v1.Group("/users")
		{
			users.GET(":id", controllers.FindUser)
			users.GET(":id/login", controllers.Login)
			users.POST("", controllers.CreateUser)
			users.PATCH(":id", controllers.UpdateUser)
			users.DELETE(":id", controllers.DeleteUser)
		}
	}

	// Swagger API route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
