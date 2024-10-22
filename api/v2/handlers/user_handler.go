package handlers

import (
	"net/http"
	"strconv"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	Conn *pgx.Conn
	Cfg  *configs.Config
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawId := c.Param("id")
	id, err := uuid.Parse(rawId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please check id and try again"})
		return
	}

	user, err := repo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)

}

func (uh *UserHandler) GetUserByUsername(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	username := c.Param("username")

	user, err := repo.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)

}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
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

	users, err := repo.GetAllUsers(c.Request.Context(),
		repository.GetAllUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)

}

func (uh *UserHandler) GetAllActiveUsers(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
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

	users, err := repo.GetActiveUsers(c.Request.Context(),
		repository.GetActiveUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)

}

func (uh *UserHandler) GetAllInActiveUsers(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
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

	users, err := repo.GetInActiveUsers(c.Request.Context(),
		repository.GetInActiveUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)

}
