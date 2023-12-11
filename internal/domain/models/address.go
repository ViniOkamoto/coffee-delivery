package models

type Address struct {
	Street       string //`form:"street" json:"street" binding:"required"`
	Number       string //`form:"number" json:"number" binding:"required"`
	Complement   string //`form:"complement" json:"complement" binding:"required"`
	Neighborhood string //`form:"neighborhood" json:"neighborhood" binding:"required"`
	City         string //`form:"city" json:"city" binding:"required"`
	State        string //`form:"state" json:"state" binding:"required"`
	Country      string //`form:"country" json:"country" binding:"required"`
	ZipCode      string //`form:"zip_code" json:"zip_code" binding:"required"`
}
