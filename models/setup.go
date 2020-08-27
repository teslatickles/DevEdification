package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// ConnectDataBase attempts to connect to mysql
// database using args passed in gorm.Open
func ConnectDataBase() {
<<<<<<< Updated upstream
	database, err := gorm.Open("sqlite3", "core.db")
=======
	// uncomment gin.SetMode for production
	// gin.SetMode(gin.ReleaseMode)

	database, err := gorm.Open("mysql", "root:1Paraprosdokian9@tcp(0.0.0.0:3306)/core?charset=utf8&parseTime=True&loc=Local")
>>>>>>> Stashed changes
	if err != nil {
		panic("Failed to connect to database!")
	}

<<<<<<< Updated upstream
=======
	databaseTest, err := gorm.Open("mysql", "root:1Paraprosdokian9@tcp(0.0.0.0:3306)/core_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

>>>>>>> Stashed changes
	dbModels := []interface{}{&Book{}, &Website{}, &VizNug{}, &User{}}
	database.AutoMigrate(dbModels...)

<<<<<<< Updated upstream
	DB = database
=======
	// this is an attempt to handle testing the api cleanly
	// but needs a lot of work/investigation/whatnot
	// help me... hehe
	if gin.ReleaseMode == "release" {
		DB = databaseTest
	} else {
		DB = databaseTest
	}
>>>>>>> Stashed changes
}
