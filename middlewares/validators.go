package middlewares

import (
	"net/http"
	"strings"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/gin-gonic/gin"
)

type MiddlewareConfig struct {
	Cfg *configs.Config
}

func (mc *MiddlewareConfig) RequireValidToken(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

	// Validate the token using ValidateJWT
	claims, err := ValidateJWT(tokenString, mc.Cfg.JWTConfig.ApiSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Next()
}
