package services

import (
	"sl-api/models"
)

func GetOilBin(id string) (models.OilBin, error) {
	bin := models.OilBin{}
	Database.First(&bin, id)

	if bin.ID == 0 {
		return bin, notFoundError
	}

	return bin, nil
}

/// Query para pesquisar oleões por endereço ou coordenadas
func SearchOilBinQuery(search string) ([]models.OilBin, error) {
	var binResults []models.OilBin

	Database.Where("coordinate_x LIKE ? OR coordinate_y LIKE ? OR address LIKE ?", search, search, search).Find(&binResults)

	if len(binResults) == 0 {
		return binResults, emptyDbError
	}

	return binResults, nil
}

func InsertOilBin(bin *models.OilBin) error {
	if Database.Create(bin).Error != nil {
		return badParamsError
	}

	return nil
}

func GetAllOilBins() ([]models.OilBin, error) {
	var allBins []models.OilBin

	Database.Find(&allBins)

	if len(allBins) == 0 {
		return allBins, emptyDbError
	}

	return allBins, nil
}

func DeleteOilBin(id string) error {
	return Database.Delete(&models.OilBin{}, id).Error
}
