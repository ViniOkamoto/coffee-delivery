package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" binding:"required" gorm:"text;not null;default:null`
	BirthDate string `json:"birth_date" binding:"required" gorm:"text;not null;default:null`
	Email     string `json:"email" binding:"required" gorm:"text;not null;default:null`
	Password  string `json:"password" binding:"required" gorm:"text;not null;default:null`
	IsAdmin   bool

	Address []Address `json:"address" binding:"false" gorm:"foreignKey:UserID"`
}

func NewUser(
	name string,
	birthdate string,

) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("Name is required")
	}
	return &User{Name: name, BirthDate: birthdate}, nil
}
