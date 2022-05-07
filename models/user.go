package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Firebase_ID int64
}
