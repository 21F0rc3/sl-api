package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/* Nomes da coluna que possui o ID dos diferentes models */
const ADDITIONAL_COMMENT_ID_ATTR_NAME = "additional_comment_id"
const BOTTLE_ID_ATTR_NAME = "bottle_id"
const DEPOSIT_ID_ATTR_NAME = "deposit_id"
const LIQUID_TYPE_ID_ATTR_NAME = "liquid_type_id"
const OIL_BIN_ID_ATTR_NAME = "oil_bin_id"
const REVIEW_STATE_ID_ATTR_NAME = "review_state_id"
const SAMPLE_LIQUID_ID_ATTR_NAME = "sample_liquid_id"
const SAMPLE_ID_ATTR_NAME = "sample_id"
const SENSOR_READING_ID_ATTR_NAME = "sensor_reading_id"
const SENSOR_TYPE_ID_ATTR_NAME = "sensor_type_id"
const SENSOR_ID_ATTR_NAME = "sensor_id"
const USER_TYPE_ID_ATTR_NAME = "user_type_id"
const USER_ID_ATTR_NAME = "user_id"
const FIREBASE_USER_UID_ATTR_NAME = "firebase_uid"

const WaterLiquidType = 1
const KitchenOilLiquidType = 2
const CarOilLiquidType = 3

const TurbiditySensorType = 1
const TemperatureSensorType = 2
const ExternalTemperatureSensorType = 3
const WaterFlowSensorType = 4
const LiquidLevelSensorType = 5
const HumiditySensorType = 6

/* Helper functions */
func addError(ctx *gin.Context, err error) {
	errErr := ctx.Error(err)
	log.Println(errErr)
}

func closeWithError(ctx *gin.Context, statusCode int, err error) {
	addError(ctx, err)
	closeWithData(ctx, statusCode, gin.H{})
}

func closeWithData(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.IndentedJSON(statusCode, gin.H{"data": data, "errors": ctx.Errors.Errors()})
}

func closeWithOk(ctx *gin.Context) {
	closeWithData(ctx, http.StatusOK, gin.H{})
}

/* Helper types */

type SampleLiquidsInput struct {
	WaterAmount  int
	OilAmount    int
	CarOilAmount int
	OthersAmount int
}

type SampleReadingsInput struct {
	LiquidTemperature   float64
	Turbidity           float64
	ExternalTemperature float64
	LiquidLevel         float64
	Humidity            float64
	WaterFlow           float64
}
