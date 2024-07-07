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
	sh := handlers.StudentHandler{Store: &db.StudentStore{DB: server.DB}, Cfg: server.Config}
	mc := middlewares.MiddleWareConfig{Cfg: server.Config}
	rh := handlers.RewardsHandler{Store: &db.RewardTransactionStore{DB: server.DB}, Cfg: server.Config}

	server.POST("/students/login/", sh.Login)

	server.POST("/students/register/", sh.RegisterStudent)
	server.GET("/students/all", mc.RequireValidToken, sh.GetAllStudents)
	server.GET("/students/all/:campus", sh.GetCampusStudents)
	server.GET("/students/find/id/:id", sh.GetStudentByID)
	server.GET("/students/find/admno/:admno", sh.GetStudentByAmno)
	server.GET("/students/registered/:admno", sh.IsStudentRegistered)
	server.GET("/students/find/username/:username", sh.GetStudentByUsername)
	server.PATCH("/students/update/:id", mc.RequireValidToken, mc.RequireSameUserOrAdmin, sh.UpdateStudent)
	server.DELETE("/students/delete/:id", mc.RequireValidToken, mc.RequireSameUserOrAdmin, sh.DeleteStudent)

	// Reward transactions
	server.POST("/rewards/award", rh.NewTransaction)
	server.GET("/rewards/leaderboard", sh.GetLeaderBoard)
	server.GET("/rewards/awards/:userid", rh.GetUserTransactions)
	server.GET("/rewards/awards/all", rh.GetAllTransactions)
	server.DELETE("/rewards/awards/:transaction", rh.DeleteRewardTransaction)
}
