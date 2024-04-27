package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    int
	ProductId int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductId"`
}
