package services

import (
	"sl-api/models"
	"sl-api/services"
)

func GetOilBin(id string) (models.OilBin, error) {
	bin := models.OilBin{}
	services.Database.First(&bin, id)

	if bin.ID == 0 {
		return bin, services.NotFoundError
	}

	return bin, nil
}

/// Query para pesquisar oleões por endereço ou coordenadas
func SearchOilBinQuery(search string) ([]models.OilBin, error) {
	var binResults []models.OilBin

	search = "%" + search + "%"

	services.Database.Where("coordinate_x::text ILIKE ? OR coordinate_y::text ILIKE ? or (coordinate_x::text || ' ' || coordinate_y::text) ilike ? OR address::text ILIKE ?", search, search, search, search).Find(&binResults)

	if len(binResults) == 0 {
		return binResults, services.EmptyDbError
	}

	return binResults, nil
}

func GetOilBinByLocation(latitude float64, longitude float64, maxRange float64) ([]models.OilBin, error) {
	var binResults []models.OilBin

	binResults, gAllErr := GetAllOilBins()
	if gAllErr != nil {
		return binResults, gAllErr
	}

	var binsNear []models.OilBin

	for _, ob := range binResults {
		if isOilBinNear(ob, latitude, longitude, maxRange) {
			binsNear = append(binsNear, ob)
		}
	}

	if len(binsNear) == 0 {
		return binsNear, services.EmptyDbError
	}

	return binsNear, nil
}

func InsertOilBin(bin *models.OilBin) error {
	if services.Database.Create(bin).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetAllOilBins() ([]models.OilBin, error) {
	var allBins []models.OilBin

	services.Database.Find(&allBins)

	if len(allBins) == 0 {
		return allBins, services.EmptyDbError
	}

	return allBins, nil
}

func DeleteOilBin(id string) error {
	return services.Database.Delete(&models.OilBin{}, id).Error
}

func isOilBinNear(oilBin models.OilBin, latitude float64, longitude float64, maxRange float64) bool {
	return maxRange < 0 || oilBin.DistanceTo(latitude, longitude) <= maxRange
}
