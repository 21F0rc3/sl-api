package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllOilBins(ctx *gin.Context) {
	allBins, err := services.GetAllOilBins()

	/* Means it has no entries on the database */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allBins)
}

func GetBin(ctx *gin.Context) {
	id := ctx.Param("oil_bin_id")

	var bin models.OilBin

	bin, getErr := services.GetOilBin(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, bin)
}

func PostBin(ctx *gin.Context) {
	var bin domain.OilBin

	bindErr := ctx.BindJSON(&bin) /* BindJSON copia o conteudo do Body do Request para o bin */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertOilBin(&bin)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, bin)
}

func DeleteOilBin(ctx *gin.Context) {
	id := ctx.Param("oil_bin_id")

	err := services.DeleteOilBin(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
