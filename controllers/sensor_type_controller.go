package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllSensorTypes(ctx *gin.Context) {
	allTypes, err := services.GetAllSensorTypes()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allTypes)
}

func GetSensorType(ctx *gin.Context) {
	var sensor_type models.SensorType

	id := ctx.Param(SENSOR_TYPE_ID_ATTR_NAME)

	sensor_type, getErr := services.GetSensorType(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sensor_type)
}

func PostSensorType(ctx *gin.Context) {
	var sensor_type models.SensorType

	bindErr := ctx.BindJSON(&sensor_type) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertSensorType(&sensor_type)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, sensor_type)
}

func DeleteSensorType(ctx *gin.Context) {
	id := ctx.Param(SENSOR_TYPE_ID_ATTR_NAME)

	err := services.DeleteSensorType(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
