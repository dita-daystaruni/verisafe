package handlers

import (
	"log"
	"net/http"

	"github.com/dita-daystaruni/verisafe/config/db"
	roleHandler "github.com/dita-daystaruni/verisafe/internal/handlers/roles"
	userhandlers "github.com/dita-daystaruni/verisafe/internal/handlers/user_handlers"
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
	suh := userhandlers.SystemUserHandler{DB: db_instance}
	server.POST("/system-users/register", suh.CreateSystemUser)
	server.GET("/system-users/all", suh.GetAllSystemUsers)
	server.GET("/system-users/find/:id", suh.GetSystemUser)
	server.PATCH("/system-users/update/:id", suh.UpdateSystemUser)
	server.DELETE("/system-users/delete/:id", suh.DeleteSystemUser)

	// Student handlers
	sh := userhandlers.StudentHandler{DB: db_instance}
	server.POST("/students/register", sh.CreateStudent)
	server.GET("/students/all", sh.GetAllStudents)
	server.GET("/students/find/:id", sh.GetStudentByID)
	server.GET("/students/details/:admno", sh.GetStudentByAdmni)
	server.PATCH("/students/update/:id", sh.UpdateStudentDetails)
	server.DELETE("/students/delete/:id", sh.DeleteStudent)

	//  Run
	server.Run(":8080")
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
