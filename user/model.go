package user

import "github.com/jinzhu/gorm"

// User es la estructura para el usuario
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
