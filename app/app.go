package app

import (
	"sl-api/controllers"
	"sl-api/services"

	"github.com/gin-gonic/gin"
)

func Run() {
	services.Setup()

	router := gin.Default()

	router.POST("/addUser", controllers.PostUser)

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
