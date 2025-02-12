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
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type RoleHandler struct {
	Conn   *pgx.Conn
	Cfg    *configs.Config
	Logger *logrus.Logger
}

func (rh *RoleHandler) RegisterRole(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var roleData repository.CreateRoleParams

	if err := c.ShouldBindJSON(&roleData); err != nil {
		return nil, errors.New("Please check your request json payload and try that again")
	}

	user, err := repo.CreateRole(c.Request.Context(), roleData)
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    roleData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    roleData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}
	return &ApiResponse{StatusCode: http.StatusCreated, Result: user}, nil
}

func (rh *RoleHandler) GetAllRoles(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
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

	roles, err := repo.ListAllRoles(c.Request.Context(),
		repository.ListAllRolesParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: roles}, nil
}

func (rh *RoleHandler) GetRoleByID(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawId := c.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, errors.New("Please provide a valid uuid")
	}

	role, err := repo.GetRoleByID(c.Request.Context(), int32(id))
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: role}, nil

}

func (rh *RoleHandler) GetRoleByName(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	name := c.Param("name")

	role, err := repo.GetRoleByNameFuzzy(c.Request.Context(), pgtype.Text{String: name})
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    name,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: role}, nil

}

func (rh *RoleHandler) UpdateRole(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var roleData repository.UpdateRoleParams

	if err := c.ShouldBindJSON(&roleData); err != nil {

		return nil, errors.New("Please check your request json payload and try that again")
	}
	role, err := repo.UpdateRole(c.Request.Context(), roleData)
	if err != nil {

		rh.Logger.WithFields(logrus.Fields{
			"payload":    roleData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    roleData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: role}, nil
}

func (rh *RoleHandler) DeleteRole(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return nil, errors.New("Please specify a valid integer id")
	}

	err = repo.DeleteRole(c.Request.Context(), int32(id))
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}

	return &ApiResponse{StatusCode: http.StatusMovedPermanently,
		Result: map[string]any{"message": "role deleted successfully"}}, nil

}

func (rh *RoleHandler) AssignPermissionToRole(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer role_id")
	}
	permissionID, err := strconv.Atoi(c.Param("permission_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer permission_id")
	}

	err = repo.AssignPermissionToRole(c.Request.Context(), repository.AssignPermissionToRoleParams{
		RoleID:       int32(roleID),
		PermissionID: int32(permissionID),
	})
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"role_id":       roleID,
			"permission_id": permissionID,
			"timestamp":     time.Now(),
			"client_ip":     c.ClientIP(),
			"user_agent":    c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"role_id":       roleID,
			"permission_id": permissionID,
			"timestamp":     time.Now(),
			"client_ip":     c.ClientIP(),
			"user_agent":    c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusCreated,
		Result: map[string]any{"message": "Permission assigned to role successfully"}}, nil
}

func (rh *RoleHandler) RemovePermissionFromRole(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer role_id")
	}
	permissionID, err := strconv.Atoi(c.Param("permission_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer permission_id")
	}

	err = repo.RemovePermissionFromRole(c.Request.Context(), repository.RemovePermissionFromRoleParams{
		RoleID:       int32(roleID),
		PermissionID: int32(permissionID),
	})
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"role_id":       roleID,
			"permission_id": permissionID,
			"timestamp":     time.Now(),
			"client_ip":     c.ClientIP(),
			"user_agent":    c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"role_id":       roleID,
			"permission_id": permissionID,
			"timestamp":     time.Now(),
			"client_ip":     c.ClientIP(),
			"user_agent":    c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK,
		Result: map[string]any{"message": "Permission removed from role successfully"}}, nil
}

func (rh *RoleHandler) ListPermissionsForRole(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer role_id")
	}

	permissions, err := repo.ListPermissionsForRole(c.Request.Context(), int32(roleID))
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"role_id":    roleID,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: permissions}, nil
}

func (rh *RoleHandler) AssignRoleToUser(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid UUID for user_id")
	}
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer role_id")
	}

	err = repo.AssignRoleToUser(c.Request.Context(), repository.AssignRoleToUserParams{
		UserID: userID,
		RoleID: int32(roleID),
	})
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"role_id":    roleID,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"role_id":    roleID,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusCreated, Result: "Role assigned to user successfully"}, nil
}

func (rh *RoleHandler) RemoveRoleFromUser(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid UUID for user_id")
	}
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid integer role_id")
	}

	err = repo.RemoveRoleFromUser(c.Request.Context(), repository.RemoveRoleFromUserParams{
		UserID: userID,
		RoleID: int32(roleID),
	})
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"role_id":    roleID,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"role_id":    roleID,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK,
		Result: map[string]any{"message": "Role removed from user successfully"}}, nil
}

func (rh *RoleHandler) ListRolesForUser(c *gin.Context) (*ApiResponse, error) {
	tx, _ := rh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return nil, errors.New("Please provide a valid UUID for user_id")
	}

	roles, err := repo.ListRolesForUser(c.Request.Context(), userID)
	if err != nil {
		rh.Logger.WithFields(logrus.Fields{
			"user_id":    userID,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: roles}, nil
}
