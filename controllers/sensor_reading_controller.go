package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllSensorReadings(ctx *gin.Context) {
	allReadings, err := services.GetAllSensorReadings()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allReadings)
}

func GetSensorReading(ctx *gin.Context) {
	var sensor_reading models.SensorReading

	id := ctx.Param(SENSOR_READING_ID_ATTR_NAME)

	sensor_reading, getErr := services.GetSensorReading(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sensor_reading)
}

func PostSensorReading(ctx *gin.Context) {
	var sensor_reading models.SensorReading

	bindErr := ctx.BindJSON(&sensor_reading) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertSensorReading(&sensor_reading)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sensor_reading)
}

func DeleteSensorReading(ctx *gin.Context) {
	id := ctx.Param(SENSOR_READING_ID_ATTR_NAME)

	err := services.DeleteSensorReading(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
