package app

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/internal/app/v2/handlers"
	"github.com/dita-daystaruni/verisafe/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(s *Server) {
	s.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	uh := handlers.UserHandler{Conn: s.Conn, Cfg: s.Config}
	ah := handlers.AuthHandler{Conn: s.Conn, Cfg: s.Config}
	mc := middlewares.MiddlewareConfig{Cfg: s.Config}

	v2 := s.Group("/v2")
	{
		v2Users := v2.Group("/users")
		{
			v2Users.POST("/register", uh.RegisterUser)
			v2Users.GET("/all", mc.RequireValidToken, uh.GetAllUsers)
			v2Users.GET("find/id/:id", uh.GetUserByID)
			v2Users.GET("find/username/:username", uh.GetUserByUsername)
			v2Users.GET("/active", uh.GetAllActiveUsers)
			v2Users.GET("/inactive", uh.GetAllInActiveUsers)
			v2Users.DELETE("/delete/:id", uh.DeleteUser)

			// User profiles
			v2Users.POST("/profile/create", uh.CreateUserProfile)
			v2Users.GET("/profile", mc.RequireValidToken, uh.GetUserProfile)
			v2Users.PATCH("/profile/update", uh.UpdateUserProfile)
		}
		v2Credentials := v2.Group("/credentials")
		{
			v2Credentials.POST("/create", uh.CreateUserCredentials)
			v2Credentials.PATCH("/change-password", uh.UpdateUserCredentials)
		}

	}

	v2Auth := v2.Group("/auth")
	{
		v2Auth.POST("/authenticate", ah.Login)
		v2Auth.GET("/logout", ah.Logout)
	}
}
