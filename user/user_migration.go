package user

import "github.com/julioc98/crudgo/db"

// Migrate migration User BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "J", Age: 20})

	// Read
	var user User
	db.First(&user, 1)               // find user with id 1
	db.First(&user, "name = ?", "J") // find user with code l1212

	// Update - update user's price to 2000
	db.Model(&user).Update("Name", "JC")

	// Delete - delete user
	// db.Delete(&user)
}
