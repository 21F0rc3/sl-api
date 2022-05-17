package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertDeposit(deposit *models.Deposit) error {
	if services.Database.Create(deposit).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetDeposit(id string) (models.Deposit, error) {
	deposit := models.Deposit{}

	services.Database.First(&deposit, id)

	if deposit.ID == 0 {
		return deposit, services.NotFoundError
	}

	return deposit, nil
}

func GetAllDeposits() ([]models.Deposit, error) {
	var allDeposits []models.Deposit

	services.Database.Find(&allDeposits)

	if len(allDeposits) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allDeposits, services.EmptyDbError
	}

	return allDeposits, nil
}

func DeleteDeposit(id string) error {
	return services.Database.Delete(&models.Deposit{}, id).Error
}
