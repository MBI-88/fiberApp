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
	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&Product{}); err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&Facture{}); err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&FactureDetail{}); err != nil {
		panic(err)
	}
	
	
	
}