package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Phone string
}

func NewUser(name string, phone string) *User {
	return &User{
		Name:  name,
		Phone: phone,
	}
}
