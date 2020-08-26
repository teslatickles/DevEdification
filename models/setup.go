package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB
var testDB *gorm.DB

func ConnectDataBase() {
	//gin.SetMode(gin.ReleaseMode)

	database, err := gorm.Open("mysql", "root:1Paraprosdokian9@tcp(127.0.0.1:3306)/core?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}

	databaseTest, err := gorm.Open("mysql", "root:1Paraprosdokian9@tcp(127.0.0.1:3306)/core_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	dbModels := []interface{}{&Book{}, &Website{}, &VizNug{}, &User{}}
	database.AutoMigrate(dbModels...)
	databaseTest.AutoMigrate(dbModels...)

	if gin.ReleaseMode == "release" {
		DB = databaseTest
	} else {
		DB = databaseTest
	}
}
