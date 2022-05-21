package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertUser(user *models.User) error {
	if services.Database.Create(user).Error != nil {
		return services.BadParamsError
	}

	return nil
}

/// Conta o numero de depositos feitos pelo o utilizador
func CountUserDeposits(firebase_uid string) (int, error) {
	type Result struct {
		count int
	}

	var result Result

	services.Database.Raw("SELECT COUNT(DEPOSITS.id) FROM Deposits JOIN BOTTLES ON bottle_id = BOTTLES.id WHERE Firebase_uid = ?", firebase_uid).Scan(&result)

	return result.count, nil
}

func GetUser(firebase_uid string) (models.User, error) {
	user := models.User{}

	// Vai buscar o primeiro user com o correspondente uid
	services.Database.First(&user, "firebase_uid = ?", firebase_uid)

	if user.ID == 0 { // Caso não encontre nenhum utilizador retorna um erro
		return user, services.NotFoundError
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var allUsers []models.User

	services.Database.Find(&allUsers)

	if len(allUsers) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allUsers, services.EmptyDbError
	}

	return allUsers, nil
}

func DeleteUser(id string) error {
	return services.Database.Delete(&models.User{}, id).Error
}
