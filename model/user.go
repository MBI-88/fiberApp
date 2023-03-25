package model

import (
	"strings"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// User ORM
type User struct {
	UserID   uint   `gorm:"primaryKey" json:"userid"`
	Email    string `gorm:"unique;size:200" json:"email"`
	Password string `gorm:"size:200" json:"password"`
	Name     string `gorm:"size:200" json:"name"`
}

// GetUsers return all users
func (u *User) GetUsers() ([]User, error) {
	var user []User
	err := DB.Find(&user).Error
	return user, err
}

// RegisterUser save user in database
func (u *User) RegisterUser() error {
	if len(u.Email) > 200 {
		return fmt.Errorf("Email to long")
	}else if !strings.Contains(u.Email,"@"){
		return fmt.Errorf("Email not contain @")
	}
	if len(u.Name) > 200 {
		return fmt.Errorf("Name to long")
	}
	if len(u.Password) > 200 {
		return fmt.Errorf("Password to long")
	}
	hashpassword, err := GenerateHasKey(u.Password)
	if err != nil {
		return fmt.Errorf("Error %s",err)
	}

	u.Password = hashpassword
	return DB.Create(u).Error
}

// Login user
func (u *User) Login() error {
	password := u.Password
	result := DB.Model(&User{}).Where("email = ?",u.Email).First(&u).Error
	if result != nil {return result}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
	return err
}

// GenerateHasKey function for generating hash
func GenerateHasKey(key string) (string, error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(key),10)
	return string(hash),err
}