package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julioc98/crudgo/util"
)

// Conn connect on SQL Data Base Gorm
func Conn() (db *gorm.DB) {

	typeDB := util.GetOSEnvironment("DB", "sqlite3")

	pathDB := util.GetOSEnvironment("DATABASE_URL", "test.db")

	db, err := gorm.Open(typeDB, pathDB)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
