package models

import (
	// gorm is imported for side-effects allowing primary_key props
	_ "github.com/jinzhu/gorm"
)

// Book TODO: add ISBN, RetailSite fields
// Book struct defining  book model
type Book struct {
	ID      uint   `json:"id" gorm:"primary_key;unique;not null"`
	Title   string `json:"title"`
	Release string `json:"release"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}
