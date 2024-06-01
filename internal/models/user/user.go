package user

import (
	"time"
)

// User represents a user model
type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	AdmissionNumber string    `json:"admission_number"`
	NationalID      string    `gorm:"uniqueIndex" json:"national_id"`
	Gender          string    `json:"gender"`
	Address         string    `json:"address"`
	Email           string    `gorm:"uniqueIndex" json:"email"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	Campus          string    `json:"campus"`
	ProfileURL      string    `json:"profile_url"`
	Password        string    `json:"password"`
	Active          bool      `json:"active"`
	DateCreated     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_updated"`
}
