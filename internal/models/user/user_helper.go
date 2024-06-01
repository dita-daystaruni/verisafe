package user

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeSave is a GORM hook that is called before saving a record
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.FirstName == "" || u.LastName == "" || u.NationalID == "" || u.Email == "" || u.Password == "" {
		return errors.New("fields must not be empty")
	}
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.AdmissionNumber = strings.TrimSpace(u.AdmissionNumber)
	u.NationalID = strings.TrimSpace(u.NationalID)
	u.Gender = strings.TrimSpace(u.Gender)
	u.Address = strings.TrimSpace(u.Address)
	u.Email = strings.TrimSpace(u.Email)
	u.Campus = strings.TrimSpace(u.Campus)
	u.ProfileURL = strings.TrimSpace(u.ProfileURL)
	u.Password = strings.TrimSpace(u.Password)
	return nil
}

// BeforeUpdate is a GORM hook that is called before updating a record
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.FirstName == "" || u.LastName == "" || u.NationalID == "" || u.Email == "" || u.Password == "" {
		return errors.New("fields must not be empty")
	}
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.AdmissionNumber = strings.TrimSpace(u.AdmissionNumber)
	u.NationalID = strings.TrimSpace(u.NationalID)
	u.Gender = strings.TrimSpace(u.Gender)
	u.Address = strings.TrimSpace(u.Address)
	u.Email = strings.TrimSpace(u.Email)
	u.Campus = strings.TrimSpace(u.Campus)
	u.ProfileURL = strings.TrimSpace(u.ProfileURL)
	u.Password = strings.TrimSpace(u.Password)
	return nil
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user from the database by its ID
func GetUserByID(db *gorm.DB, id uuid.UUID) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates a user in the database
func UpdateUser(db *gorm.DB, user *User) error {
	if user.ID == uuid.Nil {
		return errors.New("invalid user ID")
	}
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database by its ID
func DeleteUser(db *gorm.DB, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ")
	}
	if err := db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
