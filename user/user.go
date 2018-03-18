package user

// User is a Human
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Users is a slice for User
var Users []User
