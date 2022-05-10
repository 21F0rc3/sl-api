package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserType_ID int
	Firebase_ID string `gorm:"unique"`
}
