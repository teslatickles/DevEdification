package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID		 uint 	`json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"-"`
}
