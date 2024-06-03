package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateUser creates a new user in the database
func (s *SystemUser) CreateUser(db *gorm.DB) error {
	if err := db.Create(s).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user from the database by its ID
func GetSystemUserByID(db *gorm.DB, id uuid.UUID) (*SystemUser, error) {
	var user SystemUser
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database
func GetAllSystemUsers(db *gorm.DB) ([]SystemUser, error) {
	var systemUsers []SystemUser
	if err := db.Find(&systemUsers).Preload("Roles").Error; err != nil {
		return nil, err
	}
	return systemUsers, nil
}

// UpdateUser updates a user in the database
func UpdateSystemUser(db *gorm.DB, user SystemUser) error {
	if user.ID == uuid.Nil {
		return errors.New("invalid student ID")
	}
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a system user from the database
// specified by [id]
func DeleteSystemUser(db *gorm.DB, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}
	if err := db.Delete(&Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
