package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "core.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	dbModels := []interface{}{&Book{}, &Website{}, &VizNug{}, &User{}}
	database.AutoMigrate(dbModels...)

	DB = database
}
