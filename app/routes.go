package app

import (
	"sl-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	//
	//---------------------------------------------------------
	//   Additional Comment
	//---------------------------------------------------------
	//

	router.GET("/additional-comment", controllers.GetAllAdditionalComments)
	router.GET("/additional-comment/:"+controllers.ADDITIONAL_COMMENT_ID_ATTR_NAME, controllers.GetAdditionalComment)
	router.POST("/additional-comment", controllers.PostAdditionalComment)
	router.DELETE("/additional-comment/:"+controllers.ADDITIONAL_COMMENT_ID_ATTR_NAME, controllers.DeleteAdditionalComment)

	//
	//---------------------------------------------------------
	//   Bottle
	//---------------------------------------------------------
	//

	router.GET("/bottle", controllers.GetAllBottles)
	router.GET("/bottle/:"+controllers.BOTTLE_ID_ATTR_NAME, controllers.GetBottle)
	router.POST("/bottle", controllers.PostBottle)
	router.DELETE("/bottle/:"+controllers.BOTTLE_ID_ATTR_NAME, controllers.DeleteBottle)

	//
	//---------------------------------------------------------
	//   Deposit
	//---------------------------------------------------------
	//

	router.GET("/deposit", controllers.GetAllDeposits)
	router.GET("/deposit/:"+controllers.DEPOSIT_ID_ATTR_NAME, controllers.GetDeposit)
	router.POST("/deposit", controllers.PostDeposit)
	router.DELETE("/deposit/:"+controllers.DEPOSIT_ID_ATTR_NAME, controllers.DeleteDeposit)

	//
	//---------------------------------------------------------
	//   Liquid Type
	//---------------------------------------------------------
	//

	router.GET("/liquid-type", controllers.GetAllLiquidTypes)
	router.GET("/liquid-type/:"+controllers.LIQUID_TYPE_ID_ATTR_NAME, controllers.GetLiquidType)
	router.POST("/liquid-type", controllers.PostLiquidType)
	router.DELETE("/liquid-type/:"+controllers.LIQUID_TYPE_ID_ATTR_NAME, controllers.DeleteLiquidType)

	//
	//---------------------------------------------------------
	//   Oil Bin
	//---------------------------------------------------------
	//

	router.GET("/oil-bin", controllers.GetAllOilBins)
	router.GET("/oil-bin/:"+controllers.OIL_BIN_ID_ATTR_NAME, controllers.GetBin)
	router.GET("/oil-bin/search/:text_to_search", controllers.SearchOilBinQuery)
	router.POST("/oil-bin", controllers.PostBin)
	router.DELETE("/oil-bin/:"+controllers.OIL_BIN_ID_ATTR_NAME, controllers.DeleteOilBin)

	//
	//---------------------------------------------------------
	//   Review State
	//---------------------------------------------------------
	//

	router.GET("/review-state", controllers.GetAllReviewStates)
	router.GET("/review-state/:"+controllers.REVIEW_STATE_ID_ATTR_NAME, controllers.GetReviewState)
	router.POST("/review-state", controllers.PostReviewState)
	router.DELETE("/review-state/:"+controllers.REVIEW_STATE_ID_ATTR_NAME, controllers.DeleteReviewState)

	//
	//---------------------------------------------------------
	//   Sample Liquid
	//---------------------------------------------------------
	//

	router.GET("/sample-liquid", controllers.GetAllSampleLiquids)
	router.GET("/sample-liquid/:"+controllers.SAMPLE_LIQUID_ID_ATTR_NAME, controllers.GetSampleLiquid)
	router.POST("/sample-liquid", controllers.PostSampleLiquid)
	router.DELETE("/sample-liquid/:"+controllers.SAMPLE_LIQUID_ID_ATTR_NAME, controllers.DeleteSampleLiquid)

	//
	//---------------------------------------------------------
	//   Sample
	//---------------------------------------------------------
	//

	router.GET("/sample", controllers.GetAllSamples)
	router.GET("/sample/:"+controllers.SAMPLE_ID_ATTR_NAME, controllers.GetSample)
	router.POST("/sample", controllers.PostSample)
	router.DELETE("/sample/:"+controllers.SAMPLE_ID_ATTR_NAME, controllers.DeleteSample)

	//
	//---------------------------------------------------------
	//   Sample Reading
	//---------------------------------------------------------
	//

	router.GET("/sensor-reading", controllers.GetAllSensorReadings)
	router.GET("/sensor-reading/:"+controllers.SENSOR_READING_ID_ATTR_NAME, controllers.GetSensorReading)
	router.POST("/sensor-reading", controllers.PostSensorReading)
	router.DELETE("/sensor-reading/:"+controllers.SENSOR_READING_ID_ATTR_NAME, controllers.DeleteSensorReading)

	//
	//---------------------------------------------------------
	//   Sensor Type
	//---------------------------------------------------------
	//

	router.GET("/sensor-type", controllers.GetAllSensorTypes)
	router.GET("/sensor-type/:"+controllers.SENSOR_TYPE_ID_ATTR_NAME, controllers.GetSensorType)
	router.POST("/sensor-type", controllers.PostSensorType)
	router.DELETE("/sensor-type/:"+controllers.SENSOR_TYPE_ID_ATTR_NAME, controllers.DeleteSensorType)

	//
	//---------------------------------------------------------
	//   Sensor
	//---------------------------------------------------------
	//

	router.GET("/sensor", controllers.GetAllSensors)
	router.GET("/sensor/:"+controllers.SENSOR_ID_ATTR_NAME, controllers.GetSensor)
	router.POST("/sensor", controllers.PostSensor)
	router.DELETE("/sensor/:"+controllers.SENSOR_ID_ATTR_NAME, controllers.DeleteSensor)

	//
	//---------------------------------------------------------
	//   User Type
	//---------------------------------------------------------
	//

	router.GET("/user-type", controllers.GetAllUserTypes)
	router.GET("/user-type/:"+controllers.USER_TYPE_ID_ATTR_NAME, controllers.GetUserType)
	router.POST("/user-type", controllers.PostUserType)
	router.DELETE("/user-type/:"+controllers.USER_TYPE_ID_ATTR_NAME, controllers.DeleteUserType)

	//
	//---------------------------------------------------------
	//   User
	//---------------------------------------------------------
	//

	router.GET("/user", controllers.GetAllUsers)
	router.GET("/user/count-deposits/:firebase_uid", controllers.CountUserDeposits)
	router.GET("/user/:firebase_uid", controllers.GetUser)
	router.POST("/user", controllers.PostUser)
}
