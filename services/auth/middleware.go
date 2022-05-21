package fbauth

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	authorizationHeader = "Authorization"
	apiKeyHeader        = "X-API-Key"
	cronExecutedHeader  = "X-Apprentice-Cron"
	valName             = "FIREBASE_ID_TOKEN"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		authHeader := c.Request.Header.Get(authorizationHeader)
		log.Println("authHeader", authHeader)
		token := strings.Replace(authHeader, "Bearer", "", 1)
		idToken, err := firebaseClient.VerifyIDToken(c, token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
		}

		log.Println("Auth time:", time.Since(startTime))

		c.Set(valName, idToken)
		c.Next()
	}
}
