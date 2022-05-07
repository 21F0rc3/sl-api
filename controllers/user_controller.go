package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {
	var user models.User

	bindErr := ctx.BindJSON(&user) /* BindJSON copia o conteudo do Body do Request para o bin */

	/* Handle possible error of the binding */
	if bindErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, bindErr)
		return
	}

	/* Actual insertion on the Database */
	insErr := services.InsertUser(&user)

	/* Handle possible error on the insertion */
	if insErr != nil {
		closeWithError(ctx, http.StatusInternalServerError, insErr)
		return
	}

	closeWithData(ctx, http.StatusOK, user)
}
