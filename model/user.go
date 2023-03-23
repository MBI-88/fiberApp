package model

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
	return DB.Create(u).Error
}

// Login user
func (u *User) Login() (bool,error) {
	var count int64
	result := DB.Model(&User{}).Where("email = ? and password = ?", u.Email, u.Password).First(&u).Count(&count)
	return count > 0, result.Error
}
