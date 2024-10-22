package api

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/api/v2/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(s *Server) {
	s.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	uh := handlers.UserHandler{Conn: s.Conn, Cfg: s.Config}

	v2 := s.Group("/v2")
	{
		v2Users := v2.Group("/users")
		{
			v2Users.GET("/all", uh.GetAllUsers)
			v2Users.GET("find/id/:id", uh.GetUserByID)
			v2Users.GET("find/username/:username", uh.GetUserByUsername)
			v2Users.GET("/active", uh.GetAllActiveUsers)
			v2Users.GET("/inactive", uh.GetAllInActiveUsers)
		}
	}
}
