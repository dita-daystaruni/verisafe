package app

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/internal/app/v2/handlers"
	"github.com/dita-daystaruni/verisafe/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterHandlers(s *Server) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	s.Use(middlewares.LoggingMiddleware(logger))

	s.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	uh := handlers.UserHandler{Conn: s.Conn, Cfg: s.Config, Logger: logger}
	ah := handlers.AuthHandler{Conn: s.Conn, Cfg: s.Config}
	ch := handlers.CampusHandler{Conn: s.Conn, Cfg: s.Config, Logger: logger}

	v2 := s.Group("/v2")
	{
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

	// campus
	v2Campus := v2.Group("/campus")
	{
		v2Campus.POST("/register", handlers.ApiAdapter(ch.RegisterCampus))
		v2Campus.GET("/all", handlers.ApiAdapter(ch.GetAllCampuses))
		v2Campus.GET("/:id", handlers.ApiAdapter(ch.GetCampusByID))
		v2Campus.PATCH("/update/:id", handlers.ApiAdapter(ch.UpdateCampus))
		v2Campus.DELETE("/delete/:id", handlers.ApiAdapter(ch.DeleteCampus))
	}

	// users
	v2Users := v2.Group("/users")
	{
		v2Users.POST("/register", handlers.ApiAdapter(uh.RegisterUser))
		v2Users.GET("/all", handlers.ApiAdapter(uh.GetAllUsers))
		v2Users.GET("find/id/:id", handlers.ApiAdapter(uh.GetUserByID))
		v2Users.GET("find/username/:username", handlers.ApiAdapter(uh.GetUserByUsername))
		v2Users.GET("/active", handlers.ApiAdapter(uh.GetAllActiveUsers))
		v2Users.GET("/inactive", handlers.ApiAdapter(uh.GetAllInActiveUsers))
		v2Users.DELETE("/delete/:id", handlers.ApiAdapter(uh.DeleteUser))

		// User profiles
		v2Users.POST("/profile/create", uh.CreateUserProfile)
		v2Users.GET("/profile", uh.GetUserProfile)
		v2Users.PATCH("/profile/update", uh.UpdateUserProfile)
	}

}
