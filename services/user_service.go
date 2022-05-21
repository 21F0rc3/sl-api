package services

import (
	"sl-api/models"
)

func InsertUser(user *models.User) error {
	if Database.Create(user).Error != nil {
		return badParamsError
	}

	return nil
}

/// Conta o numero de depositos feitos pelo o utilizador
func CountUserDeposits(firebase_uid string) (int, error) {
	type Result struct {
		count int
	}

	var result Result

	Database.Raw("SELECT COUNT(DEPOSITS.id) FROM Deposits JOIN BOTTLES ON bottle_id = BOTTLES.id WHERE Firebase_uid = ?", firebase_uid).Scan(&result)

	return result.count, nil
}

func GetUser(firebase_uid string) (models.User, error) {
	user := models.User{}

	// Vai buscar o primeiro user com o correspondente uid
	Database.First(&user, "Firebase_UID = ?", firebase_uid)

	if user.ID == 0 { // Caso não encontre nenhum utilizador retorna um erro
		return user, notFoundError
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var allUsers []models.User

	Database.Find(&allUsers)

	if len(allUsers) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allUsers, emptyDbError
	}

	return allUsers, nil
}

func DeleteUser(id string) error {
	return Database.Delete(&models.User{}, id).Error
}
