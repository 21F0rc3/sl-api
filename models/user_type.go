package models

import "github.com/jinzhu/gorm"

type UserType struct {
	gorm.Model
	role string `gorm:"unique"`
}
