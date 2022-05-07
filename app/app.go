package app

import (
	"log"
	"os"
	"sl-api/controllers"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	services.Setup()

	router := gin.New()

	router.GET("/oil-bins", controllers.GetAllOilBins)
	router.GET("/oil-bins/:oil_bin_id", controllers.GetBin)

	router.POST("/oil-bins", controllers.PostBin)

	router.DELETE("/oil-bins/:oil_bin_id", controllers.DeleteOilBin)

	router.POST("/addUser", controllers.PostUser)

	err := router.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
