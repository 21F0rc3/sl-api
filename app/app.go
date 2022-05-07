package app

import (
	"log"
	"os"
	"sl-api/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	//services.Setup()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

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
