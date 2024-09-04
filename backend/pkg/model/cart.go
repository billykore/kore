package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Username  string
	ProductId int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductId"`
}
