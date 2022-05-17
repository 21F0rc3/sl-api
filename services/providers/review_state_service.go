package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertReviewState(review_state *models.ReviewState) error {
	if services.Database.Create(review_state).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetReviewState(id string) (models.ReviewState, error) {
	review_state := models.ReviewState{}

	services.Database.First(&review_state, id)

	if review_state.ID == 0 {
		return review_state, services.NotFoundError
	}

	return review_state, nil
}

func GetAllReviewStates() ([]models.ReviewState, error) {
	var allReviewStates []models.ReviewState

	services.Database.Find(&allReviewStates)

	if len(allReviewStates) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allReviewStates, services.EmptyDbError
	}

	return allReviewStates, nil
}

func DeleteReviewState(id string) error {
	return services.Database.Delete(&models.ReviewState{}, id).Error
}
