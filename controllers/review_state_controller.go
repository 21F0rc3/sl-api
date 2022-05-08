package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllReviewStates(ctx *gin.Context) {
	allStates, err := services.GetAllReviewStates()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allStates)
}

func GetReviewState(ctx *gin.Context) {
	var review_state models.ReviewState

	id := ctx.Param(REVIEW_STATE_ID_ATTR_NAME)

	review_state, getErr := services.GetReviewState(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, review_state)
}

func PostReviewState(ctx *gin.Context) {
	var review_state models.ReviewState

	bindErr := ctx.BindJSON(&review_state) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertReviewState(&review_state)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, review_state)
}

func DeleteReviewState(ctx *gin.Context) {
	id := ctx.Param(REVIEW_STATE_ID_ATTR_NAME)

	err := services.DeleteReviewState(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
