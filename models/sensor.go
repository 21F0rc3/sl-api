package models

import "github.com/jinzhu/gorm"

type Sensor struct {
	gorm.Model
	MacAddress   string `gorm:"unique"`
	SensorTypeID uint
}
