package model

import (
	"github.com/MBI-88/fiberApp/db"
	"gorm.io/gorm"
)



// DB connection
var DB *gorm.DB

// InitDB function 
// Migrate model
func InitDB() {
	DB = db.GetConn()
	err :=  DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}