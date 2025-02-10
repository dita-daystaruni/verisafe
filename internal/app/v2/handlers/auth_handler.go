package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/dita-daystaruni/verisafe/internal/utils"
	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	Conn   *pgx.Conn
	Cfg    *configs.Config
	Logger *logrus.Logger
}

func (ah *AuthHandler) Login(c *gin.Context) (*ApiResponse, error) {
	tx, _ := ah.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var authCreds repository.LoginInfo
	if err := c.ShouldBindJSON(&authCreds); err != nil {
		print(err.Error())
		return nil, errors.New("Please check your payload JSON and try again")
	}

	creds, err := repo.GetUserCredentials(c.Request.Context(),
		repository.GetUserCredentialsParams{
			Username:        authCreds.Username,
			Email:           authCreds.Email,
			UserID:          authCreds.UserID,
			AdmissionNumber: authCreds.AdmissionNumber,
		})
	if err != nil {
		ah.Logger.WithFields(logrus.Fields{
			"payload":    creds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return nil, errors.New("Error retrieving user please register or try again later")
	}

	err = utils.ComparePassword([]byte(*creds.Password), []byte(*authCreds.Password))
	if err != nil {
		ah.Logger.WithFields(logrus.Fields{
			"payload":    creds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return nil, errors.New("Please check your username and password and try again")
	}

	// Write to the db login time
	_, err = repo.UpdateUserCredentials(c.Request.Context(),
		repository.UpdateUserCredentialsParams{
			UserID:    creds.UserID,
			Password:  creds.Password,
			LastLogin: carbon.Now(),
		},
	)
	if err != nil {
		ah.Logger.WithFields(logrus.Fields{
			"payload":    creds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	// Fetch the user from db
	user, err := repo.GetUserByID(c.Request.Context(), creds.UserID)
	if err != nil {
		ah.Logger.WithFields(logrus.Fields{
			"payload":    creds,
			"timestamp":  time.Now(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error(err)

		return HandleDBErrors(err)
	}

	// token, err := middlewares.GenerateJWT(user, *ah.Cfg)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{
	// 		"error":   "Failed to assign token to user",
	// 		"details": err.Error(),
	// 	})
	// 	return
	//
	// }
	//
	// c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	//
	return &ApiResponse{StatusCode: http.StatusOK, Result: user}, nil
}

// Removes the authorization header from the client
func (ah *AuthHandler) Logout(c *gin.Context) (*ApiResponse, error) {

	c.Header("Authorization", "")
	return &ApiResponse{
			StatusCode: http.StatusOK,
			Result:     map[string]string{"message": "goodbye say we must, again meet we will"},
		},
		nil
}
