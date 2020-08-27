package models

import (
	// gorm imported for primary_key side-effects
	_ "github.com/jinzhu/gorm"
)
// Website struct that defines website model
type Website struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Tech    string `json:"tech"`
	Company string `json:"company"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}
