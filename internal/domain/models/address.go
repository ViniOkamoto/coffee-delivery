package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street       string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        string
	Country      string
	ZipCode      string
	UserID       uint `gorm:"references:ID"`
}
