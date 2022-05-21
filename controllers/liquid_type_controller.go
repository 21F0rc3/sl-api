package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services/providers"

	"github.com/gin-gonic/gin"
)

func GetAllLiquidTypes(ctx *gin.Context) {
	allTypes, err := services.GetAllLiquidTypes()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allTypes)
}

func GetLiquidType(ctx *gin.Context) {
	var liquid_type models.LiquidType

	id := ctx.Param(LIQUID_TYPE_ID_ATTR_NAME)

	liquid_type, getErr := services.GetLiquidType(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, liquid_type)
}

func PostLiquidType(ctx *gin.Context) {
	var liquid_type models.LiquidType

	bindErr := ctx.BindJSON(&liquid_type) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertLiquidType(&liquid_type)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, liquid_type)
}

func DeleteLiquidType(ctx *gin.Context) {
	id := ctx.Param(LIQUID_TYPE_ID_ATTR_NAME)

	err := services.DeleteLiquidType(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
