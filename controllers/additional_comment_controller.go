package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services/providers"

	"github.com/gin-gonic/gin"
)

func GetAllAdditionalComments(ctx *gin.Context) {
	allComments, err := services.GetAllAdditionalComments()

	/* Não existem registros na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allComments)
}

func GetAdditionalComment(ctx *gin.Context) {
	var comment models.AdditionalComment

	id := ctx.Param(ADDITIONAL_COMMENT_ID_ATTR_NAME)

	comment, getErr := services.GetAdditionalComment(id)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, comment)
}

func PostAdditionalComment(ctx *gin.Context) {
	var comment models.AdditionalComment

	bindErr := ctx.BindJSON(&comment) /* BindJSON copia o conteudo do Body do Request para o comment */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertAdditionalComment(&comment)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, comment)
}

func DeleteAdditionalComment(ctx *gin.Context) {
	id := ctx.Param(ADDITIONAL_COMMENT_ID_ATTR_NAME)

	err := services.DeleteAdditionalComment(id)
	if err != nil {
		closeWithError(ctx, http.StatusBadRequest, err)
	}

	closeWithData(ctx, http.StatusOK, "deleted successfully")
}
