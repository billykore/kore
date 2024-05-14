package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name        string
	Description string
}
