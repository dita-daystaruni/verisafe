package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiCors() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods("OPTIONS")

	return corsConfig
}

// Registers the various application handlers for the application
func RegisterHandlers(server *Server) {
	server.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
