package models

import (
	"github.com/jinzhu/gorm"
	"math"
)

type OilBin struct {
	gorm.Model          // { ID, CreatedAt, UpdatedAt, DeletedAt, Name }
	CoordinateX float64 `json:"coordinate_x" gorm:"coordinate_x"`
	CoordinateY float64 `json:"coordinate_y" gorm:"coordinate_y"`
	Address     string  `json:"address"`
}

func (ob OilBin) DistanceTo(latitude float64, longitude float64) float64 {
	var p = math.Pi / 180
	var c = math.Cos
	var a = 0.5 - c((latitude-ob.CoordinateX)*p)/2 +
		c(ob.CoordinateX*p)*c(latitude*p)*
			(1-c((longitude-ob.CoordinateY)*p))/2

	return 12742 * math.Asin(math.Sqrt(a))
}
