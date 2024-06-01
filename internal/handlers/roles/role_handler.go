package roleHandler

import (
	"net/http"
	"strconv"

	"github.com/dita-daystaruni/verisafe/internal/models/roles"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RoleHandler represents role handlers
type RoleHandler struct {
	DB *gorm.DB
}

// CreateRole creates a new role
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role roles.Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := roles.CreateRole(h.DB, &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// GetRoleByID retrieves a role by its ID
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	var role roles.Role
	if err := h.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// GetAllRoles retrieves all roles
func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := roles.GetAllRoles(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// UpdateRole updates a role
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role roles.Role
	if err := h.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := roles.UpdateRole(h.DB, &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// DeleteRole deletes a role by its ID
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	// Convert string to uint
	u, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user is type"})
	}

	if err := roles.DeleteRole(h.DB, uint(u)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
