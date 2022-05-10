package controllers

import (
	"net/http"
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	allUsers, err := services.GetAllUsers()

	/* Significa que não existem dados na base de dados */
	if err != nil {
		addError(ctx, err)
	}

	closeWithData(ctx, http.StatusOK, allUsers)
}

func GetUser(ctx *gin.Context) {
	firebase_uid := ctx.Param(FIREBASE_USER_UID_ATTR_NAME)

	var user models.User

	user, getErr := services.GetUser(firebase_uid)
	if getErr != nil {
		closeWithError(ctx, http.StatusNotFound, getErr)
		return
	}

	closeWithData(ctx, http.StatusOK, user)
}

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
