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
<<<<<<< Updated upstream
	r := InitGin()
	//gin.SetMode(gin.ReleaseMode)
=======
	r := initGin()
>>>>>>> Stashed changes

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

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to Vizient's Software Wizard Manual"})
	})

	return r
}
