package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dita-daystaruni/verisafe/api/auth"
	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/gin-gonic/gin"
)

type MiddleWareConfig struct {
	Cfg *configs.Config
}

func (mc *MiddleWareConfig) RequireValidToken(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	token, err := auth.VerifyToken(tokenString, mc.Cfg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Your token was not valid relogin again to continue"})
	}

	claims := token.Claims
	fmt.Printf("Claims: %v\n", claims)

	// if float64(time.Now().Unix()) > claims..(float64) {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized,
	// 		gin.H{"error": "Your token is expired please relogin to continue"})
	// }
	c.Next()
}
