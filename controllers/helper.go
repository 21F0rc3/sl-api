package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
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

/* Helper functions */
func addError(ctx *gin.Context, err error) {
	errErr := ctx.Error(err)

	log.Println(errErr)
}

func closeWithError(ctx *gin.Context, statusCode int, err error) {
	ctx.IndentedJSON(statusCode, gin.H{"data": gin.H{}, "errors": []string{err.Error()}})
}

func closeWithData(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.IndentedJSON(statusCode, gin.H{"data": data, "errors": ctx.Errors.Errors()})
}

func coalesce(str string, val string) string {
	if str != "" {
		return str
	}

	return val
}