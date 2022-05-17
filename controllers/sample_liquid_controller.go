package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services/providers"

	"github.com/gin-gonic/gin"
)

func GetAllSampleLiquids(ctx *gin.Context) {
	allSampleLiquids, err := services.GetAllSampleLiquid()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allSampleLiquids)
}

func GetSampleLiquid(ctx *gin.Context) {
	var sample_liquid models.SampleLiquid

	id := ctx.Param(SAMPLE_LIQUID_ID_ATTR_NAME)

	sample_liquid, getErr := services.GetSampleLiquid(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sample_liquid)
}

func PostSampleLiquid(ctx *gin.Context) {
	var sample_liquid models.SampleLiquid

	bindErr := ctx.BindJSON(&sample_liquid) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertSampleLiquid(&sample_liquid)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sample_liquid)
}

func DeleteSampleLiquid(ctx *gin.Context) {
	id := ctx.Param(SAMPLE_LIQUID_ID_ATTR_NAME)

	err := services.DeleteSampleLiquid(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
