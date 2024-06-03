package userhandlers

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/internal/models/user"
	"github.com/dita-daystaruni/verisafe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

// Create a student
func (sh *StudentHandler) CreateStudent(c *gin.Context) {
	var newUser user.Student
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Wrong data format, please check your password"})
		return
	}

	// Create the user
	if err := newUser.CreateStudent(sh.DB); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

// Get all students
func (sh *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := user.GetAllStudents(sh.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

// Get a student user by ID
func (sh *StudentHandler) GetStudentByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	foundStudent, err := user.GetStudentByID(sh.DB, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, foundStudent)
}

// Get a student user by admno
func (sh *StudentHandler) GetStudentByAdmni(c *gin.Context) {
	err := utils.ValidateAdmissionNumber(c.Param("admno"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid admnission "})
		return
	}

	admno := c.Param("admno")

	foundStudent, err := user.GetStudentByAdmno(sh.DB, admno)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, foundStudent)
}

// Updates a student's information specified by ID
func (sh *StudentHandler) UpdateStudentDetails(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	var u user.Student

	if err := c.ShouldBindJSON(&u); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	u.ID = id

	err = u.UpdateStudentDetails(sh.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student details updated successfully"})
}

// Delete a student specified by ID
func (sh *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please specify a valid ID"})
		return
	}

	err = user.DeleteStudent(sh.DB, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Student user deleted successfully"})
}
