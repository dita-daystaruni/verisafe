package handlers

import (
	"net/http"
	"strconv"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/dita-daystaruni/verisafe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	Conn *pgx.Conn
	Cfg  *configs.Config
}

func (uh *UserHandler) RegisterUser(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userData repository.CreateUserParams

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}

	user, err := repo.CreateUser(c.Request.Context(), userData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating user",
			"details": err.Error(),
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusCreated, user)
}

// Requires a validated jwt token claim set in context
func (uh *UserHandler) GetUserProfile(c *gin.Context) {
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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to parse user_id",
			"details": err.Error(),
		})
		return
	}

	profile, err := repo.GetUserProfile(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, profile)

}

func (uh *UserHandler) CreateUserProfile(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userData repository.CreateUserProfileParams

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}
	profile, err := repo.CreateUserProfile(c.Request.Context(), userData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating user profile",
			"details": err.Error(),
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusCreated, profile)
}

func (uh *UserHandler) UpdateUserProfile(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userData repository.UpdateUserProfileParams

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}
	profile, err := repo.UpdateUserProfile(c.Request.Context(), userData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating user profile",
			"details": err.Error(),
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusOK, profile)
}

func (uh *UserHandler) CreateUserCredentials(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userCreds repository.CreateUserCredentialsParams

	if err := c.ShouldBindJSON(&userCreds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}

	if userCreds.Password != "" {
		hashedPassword, err := utils.HashPassword(userCreds.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password please try again with different password",
			})
			return
		}

		userCreds.Password = string(hashedPassword)
	}

	creds, err := repo.CreateUserCredentials(c.Request.Context(), userCreds)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating user credentials",
			"details": err.Error(),
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusCreated, creds)
}

func (uh *UserHandler) UpdateUserCredentials(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	var userCreds repository.UpdateUserCredentialsParams

	if err := c.ShouldBindJSON(&userCreds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please check your request body and try again",
		})
		return
	}

	if userCreds.Password != "" {
		hashedPassword, err := utils.HashPassword(userCreds.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password please try again with different password",
			})
			return
		}

		userCreds.Password = string(hashedPassword)
	}

	creds, err := repo.UpdateUserCredentials(c.Request.Context(), userCreds)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating user credentials",
			"details": err.Error(),
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusCreated, creds)
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

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	tx, _ := uh.Conn.Begin(c.Request.Context())
	defer tx.Rollback(c.Request.Context())

	repo := repository.New(tx)

	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid id and try that again"})
		return
	}

	err = repo.DeleteUser(c.Request.Context(), id)
	if err != nil {

	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating user",
			"details": err.Error(),
		})
		return
	}
	if err := tx.Commit(c.Request.Context()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error committing transaction",
			"details": err.Error(),
		})
		return

	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "goodbye!"})

}
