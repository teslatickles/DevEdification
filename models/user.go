package models

import (
	// gorm is imported for primary_key prop
	_ "github.com/jinzhu/gorm"
)

// User user struct that defines user model
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"-"`
}
