package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"ufp/smartlion-api/controllers"
	"ufp/smartlion-api/services"
)

func Run() {
	services.Setup()

	router := gin.Default()

	/* Authorization required*/
	oilBin := router.Group("/oil-bin")
	oilBin.Use(services.AuthorizationRequired())
	{
		oilBin.GET("/", controllers.GetAllOilBins)
		oilBin.GET("/:id", controllers.GetBin)
		oilBin.POST("/", controllers.PostBin)
		oilBin.DELETE("/:id", controllers.DeleteOilBin)
	}

	/* Free access */
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.LoginHandler)
		auth.POST("/register", controllers.RegisterHandler)
		auth.PUT("/refresh_token", controllers.RefreshHandler)
	}

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"hello": "world!"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
