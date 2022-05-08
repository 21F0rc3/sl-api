package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllSensors(ctx *gin.Context) {
	allSensors, err := services.GetAllSensors()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allSensors)
}

func GetSensor(ctx *gin.Context) {
	var sensor models.Sensor

	id := ctx.Param(SENSOR_ID_ATTR_NAME)

	sensor, getErr := services.GetSensor(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sensor)
}

func PostSensor(ctx *gin.Context) {
	var sensor models.Sensor

	bindErr := ctx.BindJSON(&sensor) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertSensor(&sensor)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sensor)
}

func DeleteSensor(ctx *gin.Context) {
	id := ctx.Param(SENSOR_ID_ATTR_NAME)

	err := services.DeleteSensor(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
