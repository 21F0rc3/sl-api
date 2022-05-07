package models

import "github.com/jinzhu/gorm"

type AdditionalComment struct {
	gorm.Model
	SampleID int
	Comment  string `gorm:"not null"`
}
