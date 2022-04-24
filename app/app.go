package app

import (
	"github.com/gin-gonic/gin"
	"sl-api/controllers"
	"sl-api/services"
)

func Run() {
	services.Setup()

	router := gin.Default()

	router.GET("/oil-bins", controllers.GetAllOilBins)
	router.GET("/oil-bins/:oil_bin_id", controllers.GetBin)

	router.POST("/oil-bins", controllers.PostBin)

	router.DELETE("/oil-bins/:oil_bin_id", controllers.DeleteOilBin)

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
