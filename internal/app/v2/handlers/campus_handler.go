package handlers

import (
	"errors"
	"net/http"
	"strconv"
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

func (ch *CampusHandler) RegisterCampus(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var campusData repository.CreateCampusParams

	if err := c.ShouldBindJSON(&campusData); err != nil {
		return &ApiResponse{
				StatusCode: http.StatusNotAcceptable,
				Result: FormatErrorResponse("Please check your request body and try again",
					err.Error()),
			},
			nil
	}

	campus, err := repo.CreateCampus(c.Request.Context(), campusData)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"data":       campusData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    campusData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return HandleDBErrors(err)

	}
	return &ApiResponse{StatusCode: http.StatusCreated, Result: campus}, nil
}

func (ch *CampusHandler) GetAllCampuses(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return nil, errors.New("Please specify a valid numerical limit")
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return nil, errors.New("Please specify a valid numerical offset")
	}

	campuses, err := repo.GetAllCampuses(c.Request.Context(),
		repository.GetAllCampusesParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})

	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: campuses}, nil

}

func (ch *CampusHandler) GetCampusByID(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")

	id, err := uuid.Parse(rawID)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id.ID(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return nil, errors.New("Please specify a valid uuid format")

	}

	campus, err := repo.GetCampusByID(c.Request.Context(),
		id,
	)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id.ID(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: campus}, nil
}

func (ch *CampusHandler) UpdateCampus(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")

	id, err := uuid.Parse(rawID)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id.ID(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return nil, errors.New("Please specify a valid uuid format")

	}

	var campusData repository.UpdateCampusByIDParams

	if err := c.ShouldBindJSON(&campusData); err != nil {
		return &ApiResponse{
				StatusCode: http.StatusNotAcceptable,
				Result: FormatErrorResponse("Please check your request body and try again",
					err.Error()),
			},
			nil
	}

	campusData.ID = id
	campus, err := repo.UpdateCampusByID(c.Request.Context(),
		campusData,
	)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id.ID(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    campusData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return HandleDBErrors(err)

	}

	return &ApiResponse{Result: campus, StatusCode: http.StatusOK}, nil
}

func (ch *CampusHandler) DeleteCampus(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ch.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")

	id, err := uuid.Parse(rawID)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id.ID(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return nil, errors.New("Please specify a valid uuid format")

	}

	err = repo.DeleteCampus(c.Request.Context(), id)
	if err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id.ID(),
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return HandleDBErrors(err)

	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		ch.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)
		return HandleDBErrors(err)

	}

	return &ApiResponse{
			Result:     map[string]any{"message": "Campus deleted sucessfully"},
			StatusCode: http.StatusOK},
		nil
}
