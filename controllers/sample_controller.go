package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllSamples(ctx *gin.Context) {
	allSamples, err := services.GetAllSamples()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allSamples)
}

func GetSample(ctx *gin.Context) {
	var sample models.Sample

	id := ctx.Param(SAMPLE_ID_ATTR_NAME)

	sample, getErr := services.GetSample(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sample)
}

func PostSample(ctx *gin.Context) {
	var sample models.Sample

	bindErr := ctx.BindJSON(&sample) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertSample(&sample)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sample)
}

func DeleteSample(ctx *gin.Context) {
	id := ctx.Param(SAMPLE_ID_ATTR_NAME)

	err := services.DeleteSample(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
