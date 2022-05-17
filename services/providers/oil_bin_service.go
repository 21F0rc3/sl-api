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
