package authservice

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       uint
	Username string
	Email    string
	Password string
	Type     string
	Expires  string
}

type Verified struct {
	gorm.Model
	UserID uint
	Token  string
}

type Reset struct {
	gorm.Model
	UserID uint
	Token  string
}
