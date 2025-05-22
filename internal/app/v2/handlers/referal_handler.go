package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type ReferalHandler struct {
	Pool   *pgxpool.Pool
	Cfg    *configs.Config
	Logger *logrus.Logger
}

func (rh *ReferalHandler) CreateReferal(c *gin.Context) (*ApiResponse, error) {
	poolConn, err := rh.Pool.Acquire(c.Request.Context())
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    "failed to acquire pool connection",
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return nil, err
	}

	defer poolConn.Release()
	tx, _ := poolConn.Conn().Begin(c.Request.Context())
	defer func() {
		if tx != nil {
			tx.Rollback(c.Request.Context())
		}

	}()

	repo := repository.New(tx)

	var referalData repository.CreateReferalParams

	if err := c.ShouldBindJSON(&referalData); err != nil {
		return nil, errors.New("Please check your request json payload and try that again")
	}

	referal, err := repo.CreateReferal(c.Request.Context(), referalData)
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    referalData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    referalData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}
	return &ApiResponse{StatusCode: http.StatusCreated, Result: referal}, nil
}
