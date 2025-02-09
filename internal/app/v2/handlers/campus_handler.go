package handlers

import (
	"net/http"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type CampusHandler struct {
	Conn   *pgx.Conn
	Cfg    *configs.Config
	Logger *logrus.Logger
}

func (ch *CampusHandler) RegisterUser(c *gin.Context) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var campusData repository.CreateCampusParams

	if err := c.ShouldBindJSON(&campusData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}

	campus, err := repo.CreateCampus(c.Request.Context(), campusData)
	if err != nil {

		ch.Logger.WithFields(logrus.Fields{
			"data":       campusData,
			"error":      err.Error(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		})

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Something went wrong while attempting to create campus",
			"details": "See server logs for more information",
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {

		ch.Logger.WithFields(logrus.Fields{
			"data":       campusData,
			"error":      err.Error(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		})

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": "See server logs for more information",
		})
		return

	}
	c.IndentedJSON(http.StatusCreated, campus)
}

func (ch *CampusHandler) GetCampusByID(c *gin.Context) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	// repo := repository.New(tx)

	rawID := c.Param("id")

	id, err := uuid.Parse(rawID)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"data":       id.ID(),
			"error":      err.Error(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		})

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Please specify a valid campus id",
			"details": "Wrong uuid",
		})
		return

	}

	// campus, err := repo.GetCampusByID(c.Request.Context(), repository.GetCampusByIDParams{ID: id})
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Please check server for more info"})
	// 	return
	// }
	//
  c.IndentedJSON(http.StatusOK, map[string]string{"hi":"hello"})

}
