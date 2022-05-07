package app

import (
	"sl-api/controllers"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	services.Setup()

	router := gin.Default()

	router.GET("/oil-bins", controllers.GetAllOilBins)
	router.GET("/oil-bins/:oil_bin_id", controllers.GetBin)

	router.POST("/oil-bins", controllers.PostBin)

	router.DELETE("/oil-bins/:oil_bin_id", controllers.DeleteOilBin)

	router.POST("/addUser", controllers.PostUser)

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
