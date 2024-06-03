package user

import (
	"time"

	"github.com/dita-daystaruni/verisafe/internal/models/roles"
	"github.com/google/uuid"
)

// The User type is a base type that represents a typical user
type User struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserName    string       `gorm:"unique;not null" json:"username"`
	FirstName   string       `gorm:"size:30" json:"first_name"`
	LastName    string       `gorm:"size:30" json:"last_name"`
	NationalID  string       `gorm:"uniqueIndex;size:10" json:"national_id"`
	Gender      string       `json:"gender"`
	Email       string       `gorm:"uniqueIndex;size:100" json:"email"`
	Address     string       `gorm:"size:60" json:"address"`
	Password    string       `json:"password"`
	PhoneNo     string       `gorm:"size:15" json:"phone"`
	Active      bool         `json:"active"`
	ProfileURL  string       `json:"profile_url"`
	DateOfBirth time.Time    `json:"date_of_birth"`
	DateCreated time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"date_updated"`
	Roles       []roles.Role `gorm:"many2many:user_roles;" json:"roles"`
}

// The student model represents a typical student adding
// additional fields to the user type that are essential
// for student type
type Student struct {
	User
	AdmissionNumber string    `gorm:"unique;size:7" json:"admission_number"`
	DateOfAdmission time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	Course          string    `json:"course"`
	Campus          string    `json:"campus"`
}

// The SystemUser type "inherits" from the user type
// adding fields necessary for the SystemUser type
type SystemUser struct {
	User
	Department string ` json:"department"`
}
