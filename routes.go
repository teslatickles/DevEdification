package main

import (
	"github.com/DevEdification/v2/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// initRoutes invokes initGin to init the gin engine
// http framework allowing for router with baked-in goodies
// also, initializes all routes for api
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
			users.POST("", controllers.CreateUser)
			users.PATCH(":id", controllers.UpdateUser)
			users.DELETE(":id", controllers.DeleteUser)
		}
	}

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
