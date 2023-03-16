package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)


// GetConn function
func GetConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fiberApp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}