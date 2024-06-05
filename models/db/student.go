package db

import (
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
	if err := student.VerifyDetails(); err != nil {
		return nil, err
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
	if err := ss.DB.Find(&students).Error; err != nil {
		return nil, err
	}

	return &students, nil
}

// Updates a student's details in the database
func (ss *StudentStore) UpdateStudentDetails(student models.Student) (bool, error) {
	stored := models.Student{
		Username:    student.Username,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Email:       student.Email,
		Password:    student.Password,
		Campus:      student.Campus,
		DateOfBirth: student.DateOfBirth,
		Active:      student.Active,
		DateUpdated: time.Now(),
	}

	err := ss.DB.Debug().Model(&models.Student{}).Where("id = ?", student.ID).Updates(&stored).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// Deletes a student specified by id from the database
func (ss *StudentStore) DeleteStudent(id uuid.UUID) error {
	if err := ss.DB.Delete(&models.Student{}, id).Error; err != nil {
		return err
	}

	return nil
}
