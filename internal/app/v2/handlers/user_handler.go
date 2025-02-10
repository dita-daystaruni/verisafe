package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/dita-daystaruni/verisafe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Conn   *pgx.Conn
	Cfg    *configs.Config
	Logger *logrus.Logger
}

func (uh *UserHandler) RegisterUser(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userData repository.CreateUserParams

	if err := c.ShouldBindJSON(&userData); err != nil {
		return nil, errors.New("Please check your request json payload and try that again")
	}

	user, err := repo.CreateUser(c.Request.Context(), userData)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}
	return &ApiResponse{StatusCode: http.StatusCreated, Result: user}, nil
}

func (uh *UserHandler) GetUserByID(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawId := c.Param("id")
	id, err := uuid.Parse(rawId)
	if err != nil {
		return nil, errors.New("Please provide a valid uuid")
	}

	user, err := repo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: user}, nil

}

func (uh *UserHandler) GetUserByUsername(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	username := c.Param("username")

	user, err := repo.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    username,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}
	return &ApiResponse{StatusCode: http.StatusOK, Result: user}, nil
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
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

	users, err := repo.GetAllUsers(c.Request.Context(),
		repository.GetAllUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: users}, nil
}

func (uh *UserHandler) GetAllActiveUsers(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
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

	users, err := repo.GetActiveUsers(c.Request.Context(),
		repository.GetActiveUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: users}, nil

}

func (uh *UserHandler) GetAllInActiveUsers(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
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

	users, err := repo.GetInActiveUsers(c.Request.Context(),
		repository.GetInActiveUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: users}, nil
}

func (uh *UserHandler) DeleteUser(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {

		return nil, errors.New("Please specify a valid uuid format")
	}

	err = repo.DeleteUser(c.Request.Context(), id)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	if err := tx.Commit(c.Request.Context()); err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}

	return &ApiResponse{StatusCode: http.StatusMovedPermanently,
		Result: map[string]any{"message": "user deleted successfully"}}, nil

}

// Requires a validated jwt token claim set in context
func (uh *UserHandler) GetUserProfile(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	user_id, exists := c.Get("user_id")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Please check your token and retry",
		})
	}

	id, err := uuid.Parse(user_id.(string))
	if err != nil {

		return nil, errors.New("Please check your request user_id parameter and try that again")

	}

	profile, err := repo.GetUserProfile(c.Request.Context(), id)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    id,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}

	return &ApiResponse{StatusCode: http.StatusCreated, Result: profile}, nil
}

func (uh *UserHandler) CreateUserProfile(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userData repository.CreateUserProfileParams

	if err := c.ShouldBindJSON(&userData); err != nil {
		return nil, errors.New("Please check your request json payload and try that again")
	}
	profile, err := repo.CreateUserProfile(c.Request.Context(), userData)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}
	return &ApiResponse{StatusCode: http.StatusCreated, Result: profile}, nil
}

func (uh *UserHandler) UpdateUserProfile(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userData repository.UpdateUserProfileParams

	if err := c.ShouldBindJSON(&userData); err != nil {

		return nil, errors.New("Please check your request json payload and try that again")
	}
	profile, err := repo.UpdateUserProfile(c.Request.Context(), userData)
	if err != nil {

		uh.Logger.WithFields(logrus.Fields{
			"payload":    userData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}
	if err := tx.Commit(c.Request.Context()); err != nil {

		uh.Logger.WithFields(logrus.Fields{
			"payload":    userData,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: profile}, nil
}

func (uh *UserHandler) CreateUserCredentials(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userCreds repository.CreateUserCredentialsParams

	if err := c.ShouldBindJSON(&userCreds); err != nil {
		return nil, errors.New("Please check your body JSON payload and try again")
	}

	if *userCreds.Password != "" {
		hashedPassword, err := utils.HashPassword(*userCreds.Password)
		if err != nil {
			uh.Logger.WithFields(logrus.Fields{
				"payload":    userCreds,
				"timestamp":  time.Now(),
				"client_ip":  c.ClientIP(),
				"user_agent": c.Request.UserAgent(),
			}).Error(err)

			return nil, errors.New("Failed to hash your password please try again with a different password")
		}

		*userCreds.Password = string(hashedPassword)
	}

	creds, err := repo.CreateUserCredentials(c.Request.Context(), userCreds)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userCreds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		HandleDBErrors(err)
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userCreds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}
	return &ApiResponse{StatusCode: http.StatusCreated, Result: creds}, nil
}

func (uh *UserHandler) UpdateUserCredentials(c *gin.Context) (*ApiResponse, error) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userCreds repository.UpdateUserCredentialsParams

	if err := c.ShouldBindJSON(&userCreds); err != nil {

		return nil, errors.New("Please check your body JSON payload and try again")
	}

	if *userCreds.Password != "" {
		hashedPassword, err := utils.HashPassword(*userCreds.Password)
		if err != nil {

			uh.Logger.WithFields(logrus.Fields{
				"payload":    userCreds,
				"timestamp":  time.Now(),
				"client_ip":  c.ClientIP(),
				"user_agent": c.Request.UserAgent(),
			}).Error(err)

			return nil, errors.New("Failed to hash your password please try again with a different password")
		}

		*userCreds.Password = string(hashedPassword)
	}

	creds, err := repo.UpdateUserCredentials(c.Request.Context(), userCreds)
	if err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userCreds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		uh.Logger.WithFields(logrus.Fields{
			"payload":    userCreds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)

	}

	return &ApiResponse{StatusCode: http.StatusOK, Result: creds}, nil
}
