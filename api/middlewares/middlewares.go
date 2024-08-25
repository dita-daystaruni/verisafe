package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/dita-daystaruni/verisafe/api/auth"
	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type MiddleWareConfig struct {
	Cfg *configs.Config
	DB  *gorm.DB
}

func (mc *MiddleWareConfig) RequireValidToken(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	token, err := auth.VerifyToken(tokenString, mc.Cfg, mc.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Invalid token, you have. Login again to continue, you must."})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	exp, ok := claims["exp"].(float64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Fake token, this is. Trust it, we cannot. Authentication failed, it has."})
		return
	}

	if float64(time.Now().Unix()) > exp {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Expired, the token has. Renew it, you must, hmm."})
		return
	}
	c.Next()
}

func (mc *MiddleWareConfig) RequireAdmin(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	token, err := auth.VerifyToken(tokenString, mc.Cfg, mc.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Invalid token, you have. Login again to continue, you must."})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	isAdmin, ok := claims["is_admin"].(bool)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Fake token, this is. Trust it, we cannot. Authentication failed, it has."})
		return
	}

	if !isAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Denied, access is. Strong permissions, you lack."})
		return
	}

	c.Next()

}

func (mc *MiddleWareConfig) RequireSameUserOrAdmin(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	token, err := auth.VerifyToken(tokenString, mc.Cfg, mc.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Invalid token, you have. Login again to continue, you must."})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	isAdmin, ok := claims["is_admin"].(bool)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Fake token, this is. Trust it, we cannot. Authentication failed, it has."})
		return
	}

	if isAdmin {
		c.Next()
		return
	}

	if c.Param("id") == claims["user_id"].(string) || c.Param("username") == claims["username"] {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized,
		gin.H{"error": "Denied, access is. Strong permissions, you lack."})
}

func (mc *MiddleWareConfig) DeleteToken(c *gin.Context) {
	tokenString := c.GetHeader("Token")

	err := auth.DeleteToken(tokenString, mc.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": err.Error()})
		return
	}
	c.Next()
}
func (mc *MiddleWareConfig) RequireService(c *gin.Context) {
	tokenString := c.GetHeader("API-KEY")
	if tokenString != os.Getenv("API_KEY") {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "You are not authorized to perform this action"})
		return
	}

	c.Next()
}
