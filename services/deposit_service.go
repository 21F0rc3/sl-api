package services

import (
	"sl-api/models"
)

func InsertDeposit(deposit *models.Deposit) error {
	if Database.Create(deposit).Error != nil {
		return badParamsError
	}

	return nil
}

func GetDeposit(id string) (models.Deposit, error) {
	deposit := models.Deposit{}

	Database.First(&deposit, id)

	if deposit.ID == 0 {
		return deposit, notFoundError
	}

	return deposit, nil
}

func GetAllDeposits() ([]models.Deposit, error) {
	var allDeposits []models.Deposit

	Database.Find(&allDeposits)

	if len(allDeposits) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allDeposits, emptyDbError
	}

	return allDeposits, nil
}

func DeleteDeposit(id string) error {
	return Database.Delete(&models.Deposit{}, id).Error
}
