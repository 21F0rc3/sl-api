package services

import (
	"sl-api/models"
)

func InsertSampleLiquid(sample_liquid *models.SampleLiquid) error {
	if Database.Create(sample_liquid).Error != nil {
		return badParamsError
	}

	return nil
}

func GetSampleLiquid(id string) (models.SampleLiquid, error) {
	sample_liquid := models.SampleLiquid{}

	Database.First(&sample_liquid, id)

	if sample_liquid.ID == 0 {
		return sample_liquid, notFoundError
	}

	return sample_liquid, nil
}

func GetAllSampleLiquid() ([]models.SampleLiquid, error) {
	var allSampleLiquid []models.SampleLiquid

	Database.Find(&allSampleLiquid)

	if len(allSampleLiquid) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSampleLiquid, emptyDbError
	}

	return allSampleLiquid, nil
}

func DeleteSampleLiquid(id string) error {
	return Database.Delete(&models.SampleLiquid{}, id).Error
}
