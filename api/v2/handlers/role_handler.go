package handlers

import (
	"net/http"
	"strconv"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type RoleHandler struct {
	Conn *pgx.Conn
	Cfg  *configs.Config
}

func (rh *RoleHandler) CreateRole(c *gin.Context) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var role repository.CreateRoleParams
	if err := c.ShouldBindJSON(&role); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}

	r, err := repo.CreateRole(c, role)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Failed creating role",
			"details": err.Error(),
		})
		return
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, r)
}

// Lists all the roles in the database
func (rh *RoleHandler) GetAllRoles(c *gin.Context) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	println(limit, offset)

	roles, err := repo.GetAllRoles(c, repository.GetAllRolesParams{
		Limit: int32(limit), Offset: int32(offset),
	})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Failed retrieving roles",
			"details": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, roles)
}
