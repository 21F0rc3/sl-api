package controllers

import (
	"sl-api/models"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {
	var user models.User

	ctx.BindJSON(&user) /* BindJSON copia o conteudo do Body do Request para o bin */

	/* Actual insertion on the Database */
	services.InsertUser(&user)
}
