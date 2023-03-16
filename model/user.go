package model

import (
	
)

// User ORM
type User struct {
	UserID   uint   `gorm:"primaryKey" `
	Email    string `gorm:"unique;size:200"`
	Password string `gorm:"size:200"`
	Name     string `gorm:"size:200"`
}

// GetUsers return all users
func (u *User) GetUsers() ([]User, error) {
	var user []User 
	err := DB.Find(&user).Error 
	return user,err
}