package model

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	Name               string
	Description        string
	DiscountPercentage float64
	Active             bool
}
