package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertSample(sample *models.Sample) error {
	if services.Database.Create(sample).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetSample(id string) (models.Sample, error) {
	sample := models.Sample{}

	services.Database.First(&sample, id)

	if sample.ID == 0 {
		return sample, services.NotFoundError
	}

	return sample, nil
}

func GetAllSamples() ([]models.Sample, error) {
	var allSamples []models.Sample

	services.Database.Find(&allSamples)

	if len(allSamples) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSamples, services.EmptyDbError
	}

	return allSamples, nil
}

func DeleteSample(id string) error {
	return services.Database.Delete(&models.Sample{}, id).Error
}
