package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type PermissionHandler struct {
	Conn   *pgx.Conn
	Cfg    *configs.Config
	Logger *logrus.Logger
}

func IsValidPermissionName(permissionName string) bool {
	parts := strings.Split(permissionName, ":")
	if len(parts) != 2 {
		return false
	}
	return true
}

func (ph *PermissionHandler) RegisterPermission(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ph.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	permissionName := c.Param("name")
	if !IsValidPermissionName(permissionName) {
		return nil, errors.New("Please provide a valid permission name in the format 'action:resource'")
	}

	permission, err := repo.CreatePermission(c.Request.Context(), permissionName)
	if err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    permissionName,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    permissionName,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusCreated, Result: permission}, nil
}
func (ph *PermissionHandler) GetAllPermissions(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ph.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return nil, errors.New("Please provide a valid numerical limit")
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return nil, errors.New("Please provide a numerical offset")
	}

	permissions, err := repo.ListAllPermissions(c.Request.Context(),
		repository.ListAllPermissionsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: permissions}, nil
}

func (ph *PermissionHandler) GetPermissionByID(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ph.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawId := c.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, errors.New("Please provide a valid integer ID")
	}

	permission, err := repo.GetPermissionByID(c.Request.Context(), int32(id))
	if err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: permission}, nil
}

func (ph *PermissionHandler) UpdatePermission(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ph.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var permissionData repository.UpdatePermissionParams

	if err := c.ShouldBindJSON(&permissionData); err != nil {
		return nil, errors.New("Please check your request json payload and try again")
	}

	if !IsValidPermissionName(permissionData.PermissionName) {
		return nil, errors.New("Please provide a valid permission name in the format 'action:resource'")
	}

	permission, err := repo.UpdatePermission(c.Request.Context(), permissionData)
	if err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    permissionData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    permissionData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: permission}, nil
}

func (ph *PermissionHandler) DeletePermission(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ph.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return nil, errors.New("Please specify a valid integer id")
	}

	err = repo.DeletePermission(c.Request.Context(), int32(id))
	if err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		ph.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK,
		Result: map[string]any{"message": "permission deleted successfully"}}, nil
}
