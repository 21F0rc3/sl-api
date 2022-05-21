package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services/providers"

	"github.com/gin-gonic/gin"
)

func GetAllBottles(ctx *gin.Context) {
	allBottles, err := services.GetAllBottles()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allBottles)
}

func GetBottle(ctx *gin.Context) {
	var bottle models.Bottle

	id := ctx.Param(BOTTLE_ID_ATTR_NAME)

	bottle, getErr := services.GetBottle(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, bottle)
}

func PostBottle(ctx *gin.Context) {
	var bottle models.Bottle

	bindErr := ctx.BindJSON(&bottle) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertBottle(&bottle)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, bottle)
}

func DeleteBottle(ctx *gin.Context) {
	id := ctx.Param(BOTTLE_ID_ATTR_NAME)

	err := services.DeleteBottle(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
