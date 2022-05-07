package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

/* Helper functions */
func addError(ctx *gin.Context, err error) {
	errErr := ctx.Error(err)

	log.Println(errErr)
}

func closeWithError(ctx *gin.Context, statusCode int, err error) {
	ctx.IndentedJSON(statusCode, gin.H{"data": gin.H{}, "errors": []string{err.Error()}})
}

func closeWithData(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.IndentedJSON(statusCode, gin.H{"data": data, "errors": ctx.Errors.Errors()})
}
