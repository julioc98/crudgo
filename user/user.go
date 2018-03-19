package user

// User is a Human
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Users is a slice for User
var Users []User

var IDBase int
