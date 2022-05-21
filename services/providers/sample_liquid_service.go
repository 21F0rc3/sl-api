package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertSampleLiquid(sample_liquid *models.SampleLiquid) error {
	if services.Database.Create(sample_liquid).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetSampleLiquid(id string) (models.SampleLiquid, error) {
	sample_liquid := models.SampleLiquid{}

	services.Database.First(&sample_liquid, id)

	if sample_liquid.ID == 0 {
		return sample_liquid, services.NotFoundError
	}

	return sample_liquid, nil
}

func GetAllSampleLiquid() ([]models.SampleLiquid, error) {
	var allSampleLiquid []models.SampleLiquid

	services.Database.Find(&allSampleLiquid)

	if len(allSampleLiquid) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSampleLiquid, services.EmptyDbError
	}

	return allSampleLiquid, nil
}

func DeleteSampleLiquid(id string) error {
	return services.Database.Delete(&models.SampleLiquid{}, id).Error
}
