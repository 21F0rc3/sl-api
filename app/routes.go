package app

import (
	"sl-api/controllers"
	fbauth "sl-api/services/auth"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {

	//router.Use(fbauth.AuthJWT())
	//
	//---------------------------------------------------------
	//   Additional Comment
	//---------------------------------------------------------
	//
	additionalComment := router.Group("/additional-comment")
	additionalComment.Use(fbauth.AuthJWT())
	additionalComment.GET("/", controllers.GetAllAdditionalComments)
	additionalComment.GET("/:"+controllers.ADDITIONAL_COMMENT_ID_ATTR_NAME, controllers.GetAdditionalComment)
	additionalComment.POST("/", controllers.PostAdditionalComment)
	additionalComment.DELETE("/:"+controllers.ADDITIONAL_COMMENT_ID_ATTR_NAME, controllers.DeleteAdditionalComment)

	//
	//---------------------------------------------------------
	//   Bottle
	//---------------------------------------------------------
	//

	bottle := router.Group("/bottle")
	bottle.GET("/", controllers.GetAllBottles)
	bottle.GET("/:"+controllers.BOTTLE_ID_ATTR_NAME, controllers.GetBottle)
	bottle.POST("/", controllers.PostBottle)
	bottle.DELETE("/:"+controllers.BOTTLE_ID_ATTR_NAME, controllers.DeleteBottle)

	//
	//---------------------------------------------------------
	//   Deposit
	//---------------------------------------------------------
	//

	deposit := router.Group("/deposit")
	deposit.GET("/", controllers.GetAllDeposits)
	deposit.GET("/:"+controllers.DEPOSIT_ID_ATTR_NAME, controllers.GetDeposit)
	deposit.POST("/", controllers.PostDeposit)
	deposit.DELETE("/:"+controllers.DEPOSIT_ID_ATTR_NAME, controllers.DeleteDeposit)

	//
	//---------------------------------------------------------
	//   Liquid Type
	//---------------------------------------------------------
	//

	liquid_type := router.Group("/liquid-type")
	liquid_type.GET("/", controllers.GetAllLiquidTypes)
	liquid_type.GET("/:"+controllers.LIQUID_TYPE_ID_ATTR_NAME, controllers.GetLiquidType)
	liquid_type.POST("/", controllers.PostLiquidType)
	liquid_type.DELETE("/:"+controllers.LIQUID_TYPE_ID_ATTR_NAME, controllers.DeleteLiquidType)

	//
	//---------------------------------------------------------
	//   Oil Bin
	//---------------------------------------------------------
	//

	oil_bin := router.Group("/oil-bin")
	oil_bin.GET("/", controllers.GetAllOilBins)
	oil_bin.GET("/oil-bin/search/:text_to_search", controllers.SearchOilBinQuery)
	oil_bin.GET("/:"+controllers.OIL_BIN_ID_ATTR_NAME, controllers.GetBin)
	oil_bin.POST("/", controllers.PostBin)
	oil_bin.DELETE("/:"+controllers.OIL_BIN_ID_ATTR_NAME, controllers.DeleteOilBin)
	//
	//---------------------------------------------------------
	//   Review State
	//---------------------------------------------------------
	//

	review_state := router.Group("/review-state")
	review_state.GET("/", controllers.GetAllReviewStates)
	review_state.GET("/:"+controllers.REVIEW_STATE_ID_ATTR_NAME, controllers.GetReviewState)
	review_state.POST("/", controllers.PostReviewState)
	review_state.DELETE("/:"+controllers.REVIEW_STATE_ID_ATTR_NAME, controllers.DeleteReviewState)

	//
	//---------------------------------------------------------
	//   Sample Liquid
	//---------------------------------------------------------
	//

	sample_liquid := router.Group("/sample-liquid")
	sample_liquid.GET("/", controllers.GetAllSampleLiquids)
	sample_liquid.GET("/:"+controllers.SAMPLE_LIQUID_ID_ATTR_NAME, controllers.GetSampleLiquid)
	sample_liquid.POST("/", controllers.PostSampleLiquid)
	sample_liquid.DELETE("/:"+controllers.SAMPLE_LIQUID_ID_ATTR_NAME, controllers.DeleteSampleLiquid)

	//
	//---------------------------------------------------------
	//   Sample
	//---------------------------------------------------------
	//

	sample := router.Group("/sample")
	sample.GET("/", controllers.GetAllSamples)
	sample.GET("/:"+controllers.SAMPLE_ID_ATTR_NAME, controllers.GetSample)
	sample.POST("/", controllers.PostSample)
	sample.DELETE("/:"+controllers.SAMPLE_ID_ATTR_NAME, controllers.DeleteSample)

	//
	//---------------------------------------------------------
	//   Sensor Reading
	//---------------------------------------------------------
	//

	sensor_reading := router.Group("/sensor-reading")
	sensor_reading.GET("/", controllers.GetAllSensorReadings)
	sensor_reading.GET("/:"+controllers.SENSOR_READING_ID_ATTR_NAME, controllers.GetSensorReading)
	sensor_reading.POST("/", controllers.PostSensorReading)
	sensor_reading.DELETE("/:"+controllers.SENSOR_READING_ID_ATTR_NAME, controllers.DeleteSensorReading)

	//
	//---------------------------------------------------------
	//   Sensor Type
	//---------------------------------------------------------
	//

	sensor_type := router.Group("/sensor-type")
	sensor_type.GET("/", controllers.GetAllSensorTypes)
	sensor_type.GET("/:"+controllers.SENSOR_TYPE_ID_ATTR_NAME, controllers.GetSensorType)
	sensor_type.POST("/", controllers.PostSensorType)
	sensor_type.DELETE("/:"+controllers.SENSOR_TYPE_ID_ATTR_NAME, controllers.DeleteSensorType)

	//
	//---------------------------------------------------------
	//   Sensor
	//---------------------------------------------------------
	//

	sensor := router.Group("/sensor")
	sensor.GET("/", controllers.GetAllSensors)
	sensor.GET("/:"+controllers.SENSOR_ID_ATTR_NAME, controllers.GetSensor)
	sensor.POST("/", controllers.PostSensor)
	sensor.DELETE("/:"+controllers.SENSOR_ID_ATTR_NAME, controllers.DeleteSensor)

	//
	//---------------------------------------------------------
	//   User Type
	//---------------------------------------------------------
	//

	user_type := router.Group("/user-type")
	user_type.GET("/", controllers.GetAllUserTypes)
	user_type.GET("/:"+controllers.USER_TYPE_ID_ATTR_NAME, controllers.GetUserType)
	user_type.POST("/", controllers.PostUserType)
	user_type.DELETE("/:"+controllers.USER_TYPE_ID_ATTR_NAME, controllers.DeleteUserType)

	//
	//---------------------------------------------------------
	//   User
	//---------------------------------------------------------
	//

	user := router.Group("/user")
	user.GET("/", controllers.GetAllUsers)
	user.GET("/user/count-deposits/:firebase_uid", controllers.CountUserDeposits)
	user.GET("/:firebase_uid", controllers.GetUser)
	user.POST("/", controllers.PostUser)
}
