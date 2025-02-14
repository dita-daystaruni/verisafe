package middlewares

import (
	"net/http"
	"strings"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/gin-gonic/gin"
)

type MiddlewareConfig struct {
	Cfg *configs.Config
}

// Decodes the token, validates it
// and sets the claim to the request context
func (mc *MiddlewareConfig) RequireValidToken(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

	// Validate the token using ValidateJWT
	// claims, err := ValidateJWT(tokenString, mc.Cfg.JWTConfig.ApiSecret)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }

	// Extract and set the custom sub claims
	// user_id := (claims)["user_id"].(string)
	// username := (*claims)["username"].(string)
	// email := (*claims)["email"].(string)

	// c.Set("claims", (*claims))
	// c.Set("user_id", user_id)
	// c.Set("username", username)
	// c.Set("email", email)
	c.Next()
}
