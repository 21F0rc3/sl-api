package models

import "github.com/jinzhu/gorm"

type Deposit struct {
	gorm.Model
	Sample_ID int `json:"sample_id"`
	Bottle_ID int `json:"bottle_id"`
	OilBin_ID int `json:"oil_bin_id"`
}
