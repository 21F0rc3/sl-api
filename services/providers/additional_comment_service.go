package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertAdditionalComment(comment *models.AdditionalComment) error {
	if services.Database.Create(comment).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetAdditionalComment(id string) (models.AdditionalComment, error) {
	comment := models.AdditionalComment{}

	services.Database.First(&comment, id)

	if comment.ID == 0 {
		return comment, services.NotFoundError
	}

	return comment, nil
}

func GetAllAdditionalComments() ([]models.AdditionalComment, error) {
	var allComments []models.AdditionalComment

	services.Database.Find(&allComments)

	if len(allComments) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allComments, services.EmptyDbError
	}

	return allComments, nil
}

func DeleteAdditionalComment(id string) error {
	return services.Database.Delete(&models.AdditionalComment{}, id).Error
}
