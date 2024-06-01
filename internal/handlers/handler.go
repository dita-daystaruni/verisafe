package handlers

import (
	"log"
	"net/http"

	"github.com/dita-daystaruni/verisafe/config/db"
	roleHandler "github.com/dita-daystaruni/verisafe/internal/handlers/roles"
	"github.com/dita-daystaruni/verisafe/internal/handlers/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	server      = gin.Default()
	db_instance *gorm.DB
)

func init() {
	con, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to initialize database instance")
	}
	db.AutoMigrate()

	db_instance = con
}

func Serve() {
	server.GET("/ping", pong)

	// Role handling
	rh := roleHandler.RoleHandler{DB: db_instance}
	server.POST("/roles/create", rh.CreateRole)
	server.GET("/roles/all", rh.GetAllRoles)
	server.GET("/roles/find/:id", rh.GetRoleByID)
	server.PATCH("/roles/update/:id", rh.UpdateRole)
	server.DELETE("/roles/delete/:id", rh.DeleteRole)

	// User handling
	uh := users.UserHandler{DB: db_instance}
	server.POST("/users/register", uh.CreateUser)
	server.GET("/users/all", uh.GetAllUsers)
	server.GET("/users/find/:id", uh.GetUserByID)
	server.PATCH("/users/update/:id", uh.UpdateUser)
	server.DELETE("/users/delete/:id", uh.DeleteUser)

	server.Run(":8080")
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
