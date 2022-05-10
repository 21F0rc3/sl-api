package models

import "github.com/jinzhu/gorm"

type UserType struct {
	gorm.Model
	Role string `json:"type" gorm:"unique"`
}
