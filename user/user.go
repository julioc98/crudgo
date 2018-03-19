package user

import "github.com/jinzhu/gorm"

// User is a Human
type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Users is a slice for User
// var Users []User
