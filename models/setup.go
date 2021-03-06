package models

import (
	"github.com/jinzhu/gorm"
	// import gorm for mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// DB main database object to point database with migrated models
var DB *gorm.DB

// ConnectDataBase attempts to connect to mysql
// database using args passed in gorm.Open
func ConnectDataBase() {
	// uncomment gin.SetMode for production
	// gin.SetMode(gin.ReleaseMode)
	database, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/core?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}

	databaseTest, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/core_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	dbModels := []interface{}{&Book{}, &Website{}, &VizNug{}, &User{}}
	database.AutoMigrate(dbModels...)
	databaseTest.AutoMigrate(dbModels...)

	DB = databaseTest
}
