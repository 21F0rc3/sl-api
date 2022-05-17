package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services/providers"

	"github.com/gin-gonic/gin"
)

func GetAllUserTypes(ctx *gin.Context) {
	allTypes, err := services.GetAllUserTypes()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allTypes)
}

func GetUserType(ctx *gin.Context) {
	var user_type models.UserType

	id := ctx.Param(USER_TYPE_ID_ATTR_NAME)

	user_type, getErr := services.GetUserType(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, user_type)
}

func PostUserType(ctx *gin.Context) {
	var user_type models.UserType

	bindErr := ctx.BindJSON(&user_type) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertUserType(&user_type)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, user_type)
}

func DeleteUserType(ctx *gin.Context) {
	id := ctx.Param(USER_TYPE_ID_ATTR_NAME)

	err := services.DeleteUserType(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
