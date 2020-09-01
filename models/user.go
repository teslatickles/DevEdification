package models

import (
	_ "github.com/jinzhu/gorm"
)

// User user struct that defines user model
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email	 string `json:"email"`
	Role     string	`json:"role"`
}
