package roles

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Role represents a role model
type Role struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_updated"`
}

// BeforeSave is a GORM hook that is called before saving a record
func (r *Role) BeforeSave(tx *gorm.DB) (err error) {
	// Format the role name to title case before saving
	if r.Name == "" {
		return errors.New("A role must have a valid name")
	}
	r.Name = strings.Title(r.Name)
	return nil
}

// BeforeUpdate is a GORM hook that is called before updating a record
func (r *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	if r.Name == "" {
		return errors.New("A role must have a valid name")
	}
	// Update the DateUpdated field to the current time
	r.DateUpdated = time.Now()
	return nil
}

// CreateRole creates a new role in the database
func CreateRole(db *gorm.DB, role *Role) error {
	return db.Create(role).Error
}

// GetRoleByID retrieves a role from the database by its ID
func GetRoleByID(db *gorm.DB, id uint) (*Role, error) {
	var role Role
	if err := db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// GetAllRoles retrieves all roles from the database
func GetAllRoles(db *gorm.DB) ([]Role, error) {
	var roles []Role
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// UpdateRole updates a role in the database
func UpdateRole(db *gorm.DB, role *Role) error {
	return db.Save(role).Error
}

// DeleteRole deletes a role from the database by its ID
func DeleteRole(db *gorm.DB, id uint) error {
	return db.Delete(&Role{}, id).Error
}
