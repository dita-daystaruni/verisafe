package handlers

import (
	"net/http"

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

	c.IndentedJSON(http.StatusCreated, user)

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

	c.IndentedJSON(http.StatusCreated, user)

}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	users, err := repo.GetAllUsers(c.Request.Context(),
		repository.GetAllUsersParams{
			Limit:  25,
			Offset: 0,
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, users)

}

func (uh *UserHandler) GetAllActiveUsers(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	users, err := repo.GetActiveUsers(c.Request.Context(),
		repository.GetActiveUsersParams{
			Limit:  25,
			Offset: 0,
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, users)

}

func (uh *UserHandler) GetAllInActiveUsers(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	users, err := repo.GetInActiveUsers(c.Request.Context(),
		repository.GetInActiveUsersParams{
			Limit:  25,
			Offset: 0,
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, users)

}
