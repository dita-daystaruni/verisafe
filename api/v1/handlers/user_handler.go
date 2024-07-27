package handlers

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/api/auth"
	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/models"
	"github.com/dita-daystaruni/verisafe/models/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	StudentStore *db.StudentStore
	Store        *db.UserStore
	Cfg          *configs.Config
}

type LoginCreds struct {
	Username string `json:"username"`
	Admno    string `json:"admission_number"`
	Password string `json:"password"`
}

func (uh *UserHandler) Login(c *gin.Context) {
	var loginCreds LoginCreds

	if err := c.ShouldBindJSON(&loginCreds); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Invalid, the request body is. Provide data, you must."},
		)
		return
	}

	// Check if its a student login
	if loginCreds.Admno != "" {
		stud, err := uh.StudentStore.GetStudentByAdmno(loginCreds.Admno)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if ok, err := stud.ComparePassword(loginCreds.Password); err != nil || !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Please check your admission number and password"},
			)
			return
		}

		token, err := auth.GenerateToken(&stud.User, false, uh.Cfg)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,
				gin.H{"error": "Troubled, the server is. Fix it, we must."},
			)
			return
		}

		c.Header("Token", token)

		c.IndentedJSON(http.StatusOK, stud)
		return
	}

	// Otherwise its an admin login
	user, err := uh.Store.GetUserByUsername(loginCreds.Username)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"error": "The Force, not strong with this one, it is - No such user"},
		)
		return
	}
	if ok, err := user.ComparePassword(loginCreds.Password); err != nil || !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "Check username and password, you must"},
		)
		return
	}

	token, err := auth.GenerateToken(user, true, uh.Cfg)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": "Troubled, the server is. Fix it, we must."},
		)
		return
	}

	c.Header("Token", token)

	c.IndentedJSON(http.StatusOK, user)
}

func (uh *UserHandler) Logout(c *gin.Context) {
	c.Header("Token", "")
	c.IndentedJSON(http.StatusAccepted,
		gin.H{"message": "Goodbye, we must now say. Meet again, we will."},
	)

}

func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Invalid, the request body is. Provide data, you must."},
		)
		return
	}

	u, err := uh.Store.RegisterUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Error, young padawan. Failed to create user, we have.",
				"details": err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(http.StatusCreated, u)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	rawid := c.Param("id")
	id, err := uuid.Parse(rawid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to parse, the ID has. Understand it, we cannot."},
		)
		return
	}
	user, err := uh.Store.GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Find user, we could not. The Force, not strong with this search.",
				"details": err.Error(),
			},
		)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (uh *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := uh.Store.GetUserByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Find user, we could not. The Force, not strong with this search.",
				"details": err.Error(),
			},
		)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uh.Store.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to find, multiple users we expected. The Force, not strong with this search.",
				"details": err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, users)

}

func (uh *UserHandler) DeleteUserByID(c *gin.Context) {
	rawid := c.Param("id")
	id, err := uuid.Parse(rawid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to parse, the ID has. Understand it, we cannot."},
		)
		return
	}
	err = uh.Store.DeleteStudent(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to delete, the user we have. The Force, not permit this action.",
				"details": err.Error(),
			},
		)
		return
	}
	c.Status(http.StatusMovedPermanently)
}
