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
func (ss *StudentStore) UpdateStudentDetails(student models.Student) (*models.Student, error) {
	stored := models.Student{}

	if err := ss.DB.First(&stored, student.ID).Error; err != nil {
		return nil, err
	}

	stored.UserName = student.UserName
	stored.FirstName = student.FirstName
	stored.LastName = student.LastName
	stored.Active = student.Active
	stored.Campus = student.Campus
	stored.Email = student.Email
	stored.Address = student.Address
	stored.Password = student.Password
	stored.DateUpdated = time.Now()

	if err := stored.VerifyDetails(); err != nil {
		return nil, err
	}

	if err := ss.DB.Save(&stored).Error; err != nil {
		return nil, err
	}

	return &stored, nil
}

// Deletes a student specified by id from the database
func (ss *StudentStore) DeleteStudent(id uuid.UUID) error {
	if err := ss.DB.Delete(&models.Student{}, id).Error; err != nil {
		return err
	}

	return nil
}
