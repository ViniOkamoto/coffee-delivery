package models

import "fmt"

type User struct {
	Name      string `form:"name" json:"name" binding:"required"`
	BirthDate string `form:"birth_date" json:"birth_date" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	IsAdmin   bool
	Address   []Address `form:"address" json:"address" binding:"required"`
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
