package middlewares

import (
	"fmt"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start a timer to measure request duration
		start := carbon.Now()

		// Process request
		c.Next()

		// Calculate the duration and log the details
		duration := time.Since(start.StdTime())
		logger.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"duration":   fmt.Sprintf("%dμs", duration.Microseconds()),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Info("Request handled")
	}
}
