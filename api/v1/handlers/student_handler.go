package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/models"
	"github.com/dita-daystaruni/verisafe/models/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentHandler struct {
	Store *db.StudentStore
	Cfg   *configs.Config
}

func (sh *StudentHandler) EmitUserCreated(user *models.User) {
	userData, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error: Failed to marshal user data: %s\n", err.Error())
		return
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	for _, url := range sh.Cfg.EventConfig.UserCreateEvent {
		wg.Add(1) // Increment the wait group counter
		go func(url string) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(userData))
			if err != nil {
				log.Printf("Error: Failed to create request for %s: %s\n", url, err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("EVENT_API_SECRET", sh.Cfg.APISecrets.EventApiSecret)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error: Failed to send request to %s: %s\n", url, err.Error())
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Error: Received non-OK response from %s: %s\n", url, resp.Status)
			} else {
				log.Printf("Successfully sent request to %s\n", url)
			}
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func (uh *StudentHandler) GetLeaderBoard(c *gin.Context) {
	s, err := uh.Store.LeaderBoard()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, s)
}

func (uh *StudentHandler) RegisterStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := uh.Store.NewStudent(student)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uh.EmitUserCreated(&s.User)

	c.IndentedJSON(http.StatusCreated, s)
}

func (uh *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := uh.Store.GetAllStudents()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (uh *StudentHandler) GetCampusStudents(c *gin.Context) {
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

func (uh *StudentHandler) GetStudentByID(c *gin.Context) {
	rawid := c.Param("id")
	id, err := uuid.Parse(rawid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Provide a valid ID you must"})
		return
	}

	student, err := uh.Store.GetStudentByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (uh *StudentHandler) GetStudentByAmno(c *gin.Context) {
	admno := c.Param("admno")

	student, err := uh.Store.GetStudentByAdmno(admno)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (uh *StudentHandler) IsStudentRegistered(c *gin.Context) {
	admno := c.Param("admno")

	_, err := uh.Store.GetStudentByAdmno(admno)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"registered": true})
}

func (uh *StudentHandler) GetStudentByUsername(c *gin.Context) {
	username := c.Param("username")

	student, err := uh.Store.GetStudentByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (uh *StudentHandler) UpdateStudent(c *gin.Context) {
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

func (uh *StudentHandler) DeleteStudent(c *gin.Context) {
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
