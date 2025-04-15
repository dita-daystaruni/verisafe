package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic details
				logger.WithFields(logrus.Fields{
					"panic":      r,
					"timestamp":  time.Now(),
					"client_ip":  c.ClientIP(),
					"user_agent": c.Request.UserAgent(),
				}).Error("A panic occurred during request processing")

				c.JSON(http.StatusInternalServerError,
					gin.H{"error": "We crashed please retry that again"},
				)
			}
		}()
		c.Next() // Continue processing the request
	}
}
