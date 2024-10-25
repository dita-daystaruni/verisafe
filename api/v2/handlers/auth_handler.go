package handlers

import (
	"net/http"
	"time"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/dita-daystaruni/verisafe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type AuthHandler struct {
	Conn *pgx.Conn
	Cfg  *configs.Config
}

func (ah *AuthHandler) Login(c *gin.Context) {
	tx, _ := ah.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var authCreds repository.LoginInfo
	if err := c.ShouldBindJSON(&authCreds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}

	creds, err := repo.GetUserCredentials(c.Request.Context(),
		repository.GetUserCredentialsParams{
			Username: authCreds.Username,
			Email:    authCreds.Email,
			UserID:   authCreds.UserID,
		})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error retrieving user please register or try again later",
			"details": err.Error(),
		})
		return
	}

	err = utils.ComparePassword([]byte(creds.Password), []byte(authCreds.Password))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error": "Please check your username and password and try again",
		})
		return
	}

	// Write to the db login time
	_, err = repo.UpdateUserCredentials(c.Request.Context(),
		repository.UpdateUserCredentialsParams{
			UserID:    creds.UserID,
			LastLogin: time.Now(),
		},
	)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating last login time please try again later",
			"details": err.Error(),
		})
		return
	}

	// Fetch the user from db
	user, err := repo.GetUserByID(c.Request.Context(), creds.UserID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error retrieving user details please check your username and password",
			"details": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
