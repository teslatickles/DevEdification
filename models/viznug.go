package models

// gorm is imported as side-effect to set primary key in string literal
import _ "github.com/jinzhu/gorm"

// VizNug struct that defines viznug model
// GCSC is an initialism for General Computer Science Concept
type VizNug struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Tech   string `json:"tech"`
	Title  string `json:"title"`
	Author string `json:"author"`
	GCSC   bool   `json:"gcsc"`
	URL    string `json:"url"`
}
