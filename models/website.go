package models

import _ "github.com/jinzhu/gorm"

type Website struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Tech    string `json:"tech"`
	Company string `json:"company"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}
