package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `gorm:"unique"`
	Password    string
	DisplayName string `gorm:"display_name"`
}
