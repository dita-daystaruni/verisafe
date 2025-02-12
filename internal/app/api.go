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
	ah := handlers.AuthHandler{Conn: s.Conn, Cfg: s.Config, Logger: logger}
	ch := handlers.CampusHandler{Conn: s.Conn, Cfg: s.Config, Logger: logger}
	rh := handlers.RoleHandler{Conn: s.Conn, Cfg: s.Config, Logger: logger}

	v2 := s.Group("/v2")
	{
		v2Credentials := v2.Group("/credentials")
		{
			v2Credentials.POST("/create", handlers.ApiAdapter(uh.CreateUserCredentials))
			v2Credentials.PATCH("/change-password", handlers.ApiAdapter(uh.UpdateUserCredentials))
		}

	}

	v2Auth := v2.Group("/auth")
	{
		v2Auth.POST("/authenticate", handlers.ApiAdapter(ah.Login))
		v2Auth.GET("/logout", handlers.ApiAdapter(ah.Logout))
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
		v2Users.POST("/profile/create", handlers.ApiAdapter(uh.CreateUserProfile))
		v2Users.GET("/profile", handlers.ApiAdapter(uh.GetUserProfile))
		v2Users.PATCH("/profile/update", handlers.ApiAdapter(uh.UpdateUserProfile))
	}

	v2roles := v2.Group("/roles")
	{
		v2roles.POST("/create", handlers.ApiAdapter(rh.RegisterRole))
		v2roles.GET("/all", handlers.ApiAdapter(rh.GetAllRoles))
		v2roles.GET("/find/:id", handlers.ApiAdapter(rh.GetRoleByID))
		v2roles.GET("/find/name/:name", handlers.ApiAdapter(rh.GetRoleByName))
		v2roles.PATCH("/update", handlers.ApiAdapter(rh.UpdateRole))
		v2roles.DELETE("/delete/:id", handlers.ApiAdapter(rh.DeleteRole))
	}

}
