package user

import (
	"time"

	"github.com/dita-daystaruni/verisafe/internal/models/roles"
	"github.com/google/uuid"
)

// User represents a user model
type User struct {
	ID              uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserName        string       `gorm:"unique;not null" json:"username"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	AdmissionNumber string       `gorm:"unique" json:"admission_number"`
	NationalID      string       `gorm:"uniqueIndex" json:"national_id"`
	Gender          string       `json:"gender"`
	Address         string       `json:"address"`
	Email           string       `gorm:"uniqueIndex" json:"email"`
	DateOfBirth     time.Time    `json:"date_of_birth"`
	Campus          string       `json:"campus"`
	ProfileURL      string       `json:"profile_url"`
	Password        string       `json:"password"`
	Active          bool         `json:"active"`
	DateCreated     time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated     time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"date_updated"`
	Roles           []roles.Role `gorm:"many2many:user_roles;" json:"roles"`
}
