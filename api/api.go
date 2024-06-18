package api

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/api/middlewares"
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
	uh := handlers.UserHandler{Store: &db.StudentStore{DB: server.DB}, Cfg: server.Config}
	mc := middlewares.MiddleWareConfig{Cfg: server.Config}
	rh := handlers.RewardsHandler{Store: &db.RewardTransactionStore{DB: server.DB}, Cfg: server.Config}

	server.POST("/students/login/", uh.Login)

	server.POST("/students/register/", uh.RegisterStudent)
	server.GET("/students/all", mc.RequireValidToken, uh.GetAllStudents)
	server.GET("/students/all/:campus", uh.GetCampusStudents)
	server.GET("/students/find/id/:id", uh.GetStudentByID)
	server.GET("/students/find/admno/:admno", uh.GetStudentByAmno)
	server.GET("/students/registered/:admno", uh.IsStudentRegistered)
	server.GET("/students/find/username/:username", uh.GetStudentByUsername)
	server.PATCH("/students/update/:id", mc.RequireValidToken, mc.RequireSameUserOrAdmin, uh.UpdateStudent)
	server.DELETE("/students/delete/:id", mc.RequireValidToken, mc.RequireSameUserOrAdmin, uh.DeleteStudent)

	// Reward transactions
	server.POST("/rewards/award", rh.NewTransaction)
	server.GET("/rewards/leaderboard", uh.GetLeaderBoard)
	server.GET("/rewards/awards/:userid", rh.GetUserTransactions)
	server.GET("/rewards/awards/all", rh.GetAllTransactions)
	server.DELETE("/rewards/awards/:transaction", rh.DeleteRewardTransaction)
}
