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
	Store *db.StudentStore
	Cfg   *configs.Config
}

func (uh *UserHandler) Login(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please ensure you specify username and password"})
		return
	}

	s, err := uh.Store.GetStudentByUsername(student.Username)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := s.ComparePassword(student.Password); err != nil || !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Please check your username and password"})
		return
	}

	token, err := auth.GenerateToken(s.ID, false, uh.Cfg)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Token", token)

	c.IndentedJSON(http.StatusOK, s)
}

func (uh *UserHandler) RegisterStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please ensure you signup using valid json"})
		return
	}

	s, err := uh.Store.NewStudent(student)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, s)
}

func (uh *UserHandler) GetAllStudents(c *gin.Context) {
	students, err := uh.Store.GetAllStudents()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (uh *UserHandler) GetCampusStudents(c *gin.Context) {
	campus := c.Param("campus")
	if campus != "athi" && campus != "nairobi" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please ensure you use nairobi or athi"})
		return
	}
	students, err := uh.Store.GetStudentsByCampus(campus)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (uh *UserHandler) GetStudentByID(c *gin.Context) {
	rawid := c.Param("id")
	id, err := uuid.Parse(rawid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Please specify a valid ID"})
		return
	}

	student, err := uh.Store.GetStudentByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (uh *UserHandler) GetStudentByAmno(c *gin.Context) {
	admno := c.Param("admno")

	student, err := uh.Store.GetStudentByAdmno(admno)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (uh *UserHandler) GetStudentByUsername(c *gin.Context) {
	username := c.Param("username")

	student, err := uh.Store.GetStudentByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (uh *UserHandler) UpdateStudent(c *gin.Context) {
	rawid := c.Param("id")
	id, err := uuid.Parse(rawid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Please specify a valid ID"})
		return
	}

	var student models.Student
	student.ID = id
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please ensure you signup using valid json"})
		return
	}

	_, err = uh.Store.UpdateStudentDetails(student)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student details updated successfully"})
}

func (uh *UserHandler) DeleteStudent(c *gin.Context) {
	rawid := c.Param("id")
	id, err := uuid.Parse(rawid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Please specify a valid ID"})
		return
	}

	err = uh.Store.DeleteStudent(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
