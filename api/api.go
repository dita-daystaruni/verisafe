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

	uh := handlers.UserHandler{
		Store:        &db.UserStore{DB: server.DB},
		Cfg:          server.Config,
		StudentStore: &db.StudentStore{DB: server.DB},
	}
	sh := handlers.StudentHandler{Store: &db.StudentStore{DB: server.DB}, Cfg: server.Config}
	mc := middlewares.MiddleWareConfig{Cfg: server.Config, DB: server.DB}
	rh := handlers.RewardsHandler{Store: &db.RewardTransactionStore{DB: server.DB}, Cfg: server.Config}

	// User handler
	server.POST("/users/login/", uh.Login)
	server.GET("/users/logout/", uh.Logout)
	server.POST("users/register/", mc.RequireAdmin, uh.RegisterUser)
	server.GET("users/find/:id", mc.RequireAdmin, uh.GetUserByID)
	server.GET("users/find/username/:username", mc.RequireAdmin, uh.GetUserByUsername)
	server.GET("users/all", mc.RequireAdmin, uh.GetAllUsers)
	server.DELETE("users/delete/:id", mc.RequireAdmin, uh.DeleteUserByID)

	// Student handlers
	server.POST("/students/register/", sh.RegisterStudent)
	server.GET("/students/all", mc.RequireValidToken, mc.RequireAdmin, sh.GetAllStudents)
	server.GET("/students/all/:campus", mc.RequireAdmin, sh.GetCampusStudents)
	server.GET("/students/find/id/:id", mc.RequireSameUserOrAdmin, sh.GetStudentByID)
	server.GET("/students/find/admno/:admno", mc.RequireAdmin, sh.GetStudentByAmno)
	server.GET("/students/registered/:admno", sh.IsStudentRegistered)
	server.GET("/students/find/username/:username", mc.RequireSameUserOrAdmin, sh.GetStudentByUsername)
	server.PATCH("/students/update/:id", mc.RequireValidToken, mc.RequireSameUserOrAdmin, sh.UpdateStudent)
	server.DELETE("/students/delete/:id", mc.RequireValidToken, mc.RequireSameUserOrAdmin, sh.DeleteStudent)

	// Reward transactions
	server.POST("/rewards/award", mc.RequireAdmin, rh.NewTransaction)
	server.GET("/rewards/leaderboard", mc.RequireValidToken, sh.GetLeaderBoard)
	server.GET("/rewards/awards/:userid", mc.RequireSameUserOrAdmin, rh.GetUserTransactions)
	server.GET("/rewards/awards/all", mc.RequireValidToken, rh.GetAllTransactions)
	server.DELETE("/rewards/awards/:transaction", mc.RequireValidToken, rh.DeleteRewardTransaction)
}
