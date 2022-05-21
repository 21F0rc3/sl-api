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

	services.Database.Where("coordinate_x LIKE ? OR coordinate_y LIKE ? OR address LIKE ?", search, search, search).Find(&binResults)

	if len(binResults) == 0 {
		return binResults, services.EmptyDbError
	}

	return binResults, nil
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
