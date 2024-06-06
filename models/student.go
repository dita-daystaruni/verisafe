package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Student struct {
	ID                 uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Username           string              `gorm:"uniqueIndex;not null" json:"username"`
	FirstName          string              `json:"first_name"`
	LastName           string              `json:"last_name"`
	AdmissionNumber    string              `gorm:"uniqueIndex" json:"admission_number"`
	NationalID         string              `gorm:"uniqueIndex" json:"national_id"`
	Gender             string              `json:"gender"`
	Address            string              `json:"address"`
	Email              string              `gorm:"uniqueIndex" json:"email"`
	DateOfBirth        time.Time           `json:"date_of_birth"`
	Campus             string              `json:"campus"`
	ProfileURL         string              `json:"profile_url"`
	Password           string              `json:"password"`
	Active             bool                `gorm:"default:true" json:"active"`
	VibePoints         int                 `gorm:"default:0" json:"vibe_points"`
	PointsTransactions []RewardTransaction `json:"point_transactions"`
	DateCreated        time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated        time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"date_updated"`
}

// Verifies the username is in the format 00-0000
func VerifyAdmno(admno string) bool {
	pattern := `^[0-9]{2}-[0-9]{4}$`
	match, _ := regexp.MatchString(pattern, admno)
	return match
}

// HashPassword hashes the password from plaintext ready for storage
func (u *Student) HashPassword() error {
	pwd := []byte(u.Password)

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return nil
}

// The ComparePassword function compares the hash if it maches with
// the plaintext password
func (u *Student) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *Student) VerifyDetails() error {
	// Username
	if len(u.Username) < 3 {
		return errors.New("Please try a username with more than 3 characters")
	}

	// Names
	if len(u.FirstName) < 3 || len(u.LastName) < 3 {
		return errors.New("Please provide your valid names")
	}

	if len(u.AdmissionNumber) < 7 && !VerifyAdmno(u.AdmissionNumber) {
		return errors.New("Please provide a valid admission number in the 00-0000 format")
	}

	if u.Address == "" {
		return errors.New("Please specify your address")
	}

	if len(u.Email) < 5 {
		return errors.New("Please provide a valid email address")
	}

	if len(u.Password) < 6 {
		return errors.New("Please provide a valid password")
	}

	if u.Gender != "male" && u.Gender != "female" {
		return errors.New("Please specify your gender to be either male or female")
	}

	if u.Campus != "athi" && u.Campus != "nairobi" {
		return errors.New("Please specify a valid campus athi or nairobi")
	}

	return nil
}

// hooks
func (u *Student) BeforeCreate(tx *gorm.DB) (err error) {
	return u.HashPassword()
}

func (u *Student) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.HashPassword()
}
