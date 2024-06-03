package user

import (
	"errors"

	"github.com/dita-daystaruni/verisafe/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Student) BeforeSave(tx *gorm.DB) (err error) {
	return utils.ValidateAdmissionNumber(s.AdmissionNumber)
}

// CreateUser creates a new user in the database
func (s *Student) CreateStudent(db *gorm.DB) error {
	if err := db.Create(s).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user from the database by its ID
func GetStudentByID(db *gorm.DB, id uuid.UUID) (*Student, error) {
	var found Student
	if err := db.First(&found, id).Error; err != nil {
		return nil, err
	}
	return &found, nil
}

// GetUserByID retrieves a user from the database by its ID
func GetStudentByAdmno(db *gorm.DB, admno string) (*Student, error) {
	var found Student
	if err := db.Where("admission_number = ?", admno).First(&found).Error; err != nil {
		return nil, err
	}
	return &found, nil
}

// GetAllUsers retrieves all users from the database
func GetAllStudents(db *gorm.DB) ([]Student, error) {
	var students []Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// UpdateUser updates a user in the database
func (s *Student) UpdateStudentDetails(db *gorm.DB) error {
	if s.ID == uuid.Nil {
		return errors.New("invalid student ID")
	}
	if err := db.Save(s).Error; err != nil {
		return err
	}
	return nil
}

func DeleteStudent(db *gorm.DB, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ")
	}
	if err := db.Delete(&Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
