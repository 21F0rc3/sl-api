package services

import (
	"sl-api/models"
)

func InsertSample(sample *models.Sample) error {
	if Database.Create(sample).Error != nil {
		return badParamsError
	}

	return nil
}

func GetSample(id string) (models.Sample, error) {
	sample := models.Sample{}

	Database.First(&sample, id)

	if sample.ID == 0 {
		return sample, notFoundError
	}

	return sample, nil
}

func GetAllSamples() ([]models.Sample, error) {
	var allSamples []models.Sample

	Database.Find(&allSamples)

	if len(allSamples) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSamples, emptyDbError
	}

	return allSamples, nil
}

func DeleteSample(id string) error {
	return Database.Delete(&models.Sample{}, id).Error
}
