package models

import "github.com/jinzhu/gorm"

type LiquidType struct {
	gorm.Model
	Liquid string `gorm:"unique"`
}
