package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services/providers"

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

// PostLabSample
// Description:
// 		This endpoint serves the purpose of handling incoming
//		test samples that will later converge into the AI. It
//		shall not be used in production.
//
// Expects a POST request body with
//	{
//		LiquidTemperature,
//		Turbidity,
//		ExternalTemperature,
//		LiquidLevel,
//		...
//	}
func PostLabSample(ctx *gin.Context) {
	var input struct {
		readings SampleReadingsInput
		liquids  SampleLiquidsInput
	}

	bindErr := ctx.BindJSON(&input) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Insert Sample entry */
	sample, sErr := services.InsertSample()
	if sErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, sErr)
		return
	}

	/* Handle liquid quantities */
	liquids := input.liquids
	lErr := services.AssociateLiquids(sample.ID, liquids)
	if lErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, lErr)
		return
	}

	/* Handle sensor readings */
	srErr := services.AssociateSensorReadings(sample.ID, input.readings)
	if srErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, srErr)
		return
	}

	closeWithOk(ctx)
}

func DeleteSample(ctx *gin.Context) {
	id := ctx.Param(SAMPLE_ID_ATTR_NAME)

	err := services.DeleteSample(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
