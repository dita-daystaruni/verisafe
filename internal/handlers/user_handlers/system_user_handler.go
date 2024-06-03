package userhandlers

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/internal/models/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SystemUserHandler struct {
	DB *gorm.DB
}

// Create a system user
func (suh *SystemUserHandler) CreateSystemUser(c *gin.Context) {
	var newUser user.SystemUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Wrong data format, please check your password"})
		return
	}

	// Create the user
	if err := newUser.CreateUser(suh.DB); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

// Get all system users
func (suh *SystemUserHandler) GetAllSystemUsers(c *gin.Context) {
	systemUsers, err := user.GetAllSystemUsers(suh.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, systemUsers)
}

// Get s system user by ID
func (suh *SystemUserHandler) GetSystemUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	foundUser, err := user.GetSystemUserByID(suh.DB, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, foundUser)
}

// Get s system user by ID
func (suh *SystemUserHandler) UpdateSystemUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	var u user.SystemUser

	if err := c.ShouldBindJSON(&u); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	u.ID = id

	err = user.UpdateSystemUser(suh.DB, u)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User details updated successfully"})
}

// Get s system user by ID
func (suh *SystemUserHandler) DeleteSystemUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	err = user.DeleteSystemUser(suh.DB, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "System user deleted successfully"})
}
