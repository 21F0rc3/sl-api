package services

import (
	"fmt"
	"ufp/smartlion-api/domain"
)

/* Error messages */
var (
	notFoundError  = fmt.Errorf("oil bin not found")
	emptyDbError   = fmt.Errorf("there are no bins loaded")
	badParamsError = fmt.Errorf("not enough or invalid parameters")
)

func GetOilBin(id string) (domain.OilBin, error) {
	bin := domain.OilBin{}
	Database.First(&bin, id)

	if bin.ID == 0 {
		return bin, notFoundError
	}

	return bin, nil
}

func InsertOilBin(bin *domain.OilBin) error {
	if Database.Create(bin).Error != nil {
		return badParamsError
	}

	return nil
}

func GetAllOilBins() ([]domain.OilBin, error) {
	var allBins []domain.OilBin

	Database.Find(&allBins)

	if len(allBins) == 0 {
		return allBins, emptyDbError
	}

	return allBins, nil
}

func DeleteOilBin(id string) error {
	return Database.Delete(&domain.OilBin{}, id).Error
}
