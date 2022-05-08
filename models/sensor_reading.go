package models

import "github.com/jinzhu/gorm"

type SensorReading struct {
	gorm.Model
	SensorID int
	Value    float32
	SampleID int
}
