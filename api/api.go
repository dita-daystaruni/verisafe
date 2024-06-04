package api

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/api/v1/handlers"
	"github.com/dita-daystaruni/verisafe/models/db"
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

	// Student CRUD operations
	uh := handlers.UserHandler{Store: &db.StudentStore{DB: server.DB}}

	server.POST("/students/register/", uh.RegisterStudent)
	server.GET("/students/all", uh.GetAllStudents)
	server.GET("/students/all/:campus", uh.GetCampusStudents)
	server.GET("/students/find/id/:id", uh.GetStudentByID)
	server.GET("/students/find/admno/:admno", uh.GetStudentByAmno)
	server.PATCH("/students/update/:id", uh.UpdateStudent)
	server.Delete("/students/delete/:id", uh.DeleteStudent)
}
