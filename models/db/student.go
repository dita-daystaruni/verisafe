package db

import (
	"errors"
	"strings"
	"time"

	"github.com/dita-daystaruni/verisafe/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentStore struct {
	DB *gorm.DB
}

// Saves a student onto the database
func (ss *StudentStore) NewStudent(student models.Student) (*models.Student, error) {
	if err := student.Validate(); err != nil {
		return nil, err
	}
	if s, _ := ss.GetStudentByUsername(student.Username); s != nil {
		return nil, errors.New("That username is already taken please try another one")
	}
	student.DateCreated = time.Now()
	student.DateUpdated = time.Now()

	if err := ss.DB.Save(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

// Retrieves a student specified by ID
func (ss *StudentStore) GetStudentByID(id uuid.UUID) (*models.Student, error) {
	var student models.Student
	if err := ss.DB.First(&student, id).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

// Retrieves a student specified by username
func (ss *StudentStore) GetStudentByUsername(username string) (*models.Student, error) {
	var student models.Student
	if err := ss.DB.Where("username = ? ", username).First(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

// Retrieves a student specified by admno
func (ss *StudentStore) GetStudentByAdmno(admno string) (*models.Student, error) {
	var student models.Student
	if err := ss.DB.Where("admission_number = ? ", admno).First(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

// Retrieves all students  from a specified campus
func (ss *StudentStore) GetStudentsByCampus(campus string) (*[]models.Student, error) {
	var students []models.Student
	if err := ss.DB.Where("campus = ? ", campus).Find(&students).Error; err != nil {
		return nil, err
	}

	return &students, nil
}

// Retrieves all students stored
func (ss *StudentStore) GetAllStudents() (*[]models.Student, error) {
	var students []models.Student
	if err := ss.DB.Select(
		"id", "admission_number", "username",
		"first_name", "last_name", "national_id",
		"vibe_points", "address", "date_of_birth",
		"gender", "date_created", "email", "campus",
	).Find(&students).Error; err != nil {
		return nil, err
	}

	return &students, nil
}

// Updates a student's details in the database
func (ss *StudentStore) UpdateStudentDetails(student models.Student) (bool, error) {
	stored := models.Student{}
	stored.Active = student.Active
	stored.Username = student.Username
	stored.Password = student.Password
	stored.DateUpdated = time.Now()
	stored.BeforeUpdate(ss.DB)

	err := ss.DB.Debug().Model(&models.Student{}).Where("id = ?", student.ID).Updates(&stored).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// Deletes a student specified by id from the database
func (ss *StudentStore) DeleteStudent(id uuid.UUID) error {
	if err := ss.DB.Where("student_id = ?", id).
		Delete(&models.RewardTransaction{}).Error; err != nil {
		return err
	}

	if err := ss.DB.Delete(&models.Student{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (ss *StudentStore) LeaderBoard() (*[]models.Student, error) {
	var students []models.Student

	err := ss.DB.Select("username, first_name, vibe_points, profile_url, gender").Order("vibe_points DESC").Limit(25).Find(&students).Error
	if err != nil {
		return nil, err
	}

	return &students, nil
}

func (ss *StudentStore) UpdateProfilePicture(student models.Student, filename string, host string) (*models.Student, error) {

	stored := student
	stored.Password = ""
	stored.ProfileURL = host + "/" + "uploads/profiles/" + strings.ReplaceAll(student.ID.String()+filename, "/", "-")
	err := ss.DB.Debug().Model(&models.Student{}).Where("id = ?", stored.ID).Update("profile_url", stored.ProfileURL).First(&stored).Error
	if err != nil {
		return nil, err
	}

	return &stored, nil

}
