package main

import (
	"github.com/DevEdification/v2/models"
)

func main() {
	models.ConnectDataBase()
	initRoutes()
}
