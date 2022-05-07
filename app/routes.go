package app

import (
	"sl-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	//
	//---------------------------------------------------------
	//   Oil Bin
	//---------------------------------------------------------
	//

	router.GET("/oil-bins", controllers.GetAllOilBins)
	router.GET("/oil-bins/:oil_bin_id", controllers.GetBin)
	router.POST("/oil-bins", controllers.PostBin)
	router.DELETE("/oil-bins/:oil_bin_id", controllers.DeleteOilBin)

	//
	//---------------------------------------------------------
	//   User
	//---------------------------------------------------------
	//

	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/addUser", controllers.PostUser)
}
