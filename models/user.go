package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserType_ID  int    `json:"user_type_id"`
	Firebase_UID string `json:"firebase_uid" gorm:"unique"`
}
