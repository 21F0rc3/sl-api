package app

import (
	"log"
	"os"
	"sl-api/services"
	fbauth "sl-api/services/auth"

	"github.com/gin-gonic/gin"
)

func Run() {
	fbErr := fbauth.InitAuth()
	if fbErr != nil {
		log.Fatalln("Failed to initialize firebase auth", fbErr)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	services.Setup()

	router := gin.Default()

	SetRoutes(router)

	// RUN
	err := router.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
