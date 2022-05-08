package models

import "github.com/jinzhu/gorm"

type Deposit struct {
	gorm.Model
	SampleID int
	BottleID int
	OilBinID int
}
