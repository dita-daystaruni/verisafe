package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Username    string    `gorm:"uniqueIndex;not null" json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	ProfileURL  string    `json:"profile_url"`
	NationalID  string    `gorm:"uniqueIndex" json:"national_id"`
	Email       string    `gorm:"uniqueIndex" json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Active      bool      `gorm:"default:true" json:"active"`
	Password    string    `json:"password"`
	DateCreated time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_updated"`
}

// Hooks for before saving

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.HashPassword()
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.DateUpdated = time.Now()
	return u.HashPassword()
}

// HashPassword hashes the password from plaintext ready for storage
func (u *User) HashPassword() error {
	print(u.Password)
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
func (u *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

// Validates the user's information ensuring that it meets
// some certain non-empty criteria
func (u *User) Validate() error {
	// Username
	if len(u.Username) < 3 {
		return errors.New("Please try a username with more than 3 characters")
	}

	// Names
	if len(u.FirstName) < 3 || len(u.LastName) < 3 {
		return errors.New("Please provide your valid names")
	}

	if u.Address == "" {
		return errors.New("Please specify your address")
	}

	if !isValidEmail(u.Email) {
		return errors.New("Please provide a valid email address")
	}

	if len(u.Password) < 6 {
		return errors.New("Please provide a valid password")
	}

	if u.Gender != "male" && u.Gender != "female" {
		return errors.New("Please specify your gender to be either male or female")
	}
	return nil
}

func isValidEmail(email string) bool {
	// Regex for simple email validation (not fully compliant with RFC 5322)
	// This regex checks for basic format "user@example.com"
	// For production, consider using a more comprehensive regex or email validation library
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}
