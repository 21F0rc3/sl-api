package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	user_type_id int
	Firebase_ID  string `gorm:"unique"`
}
