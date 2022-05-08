package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllDeposits(ctx *gin.Context) {
	allDeposits, err := services.GetAllDeposits()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allDeposits)
}

func GetDeposit(ctx *gin.Context) {
	var deposit models.Deposit

	id := ctx.Param(DEPOSIT_ID_ATTR_NAME)

	deposit, getErr := services.GetDeposit(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, deposit)
}

func PostDeposit(ctx *gin.Context) {
	var deposit models.Deposit

	bindErr := ctx.BindJSON(&deposit) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertDeposit(&deposit)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, deposit)
}

func DeleteDeposit(ctx *gin.Context) {
	id := ctx.Param(DEPOSIT_ID_ATTR_NAME)

	err := services.DeleteDeposit(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
