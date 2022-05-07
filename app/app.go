package app

import (
	"log"
	"os"
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

	Routes(router)

	// RUN
	err := router.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
