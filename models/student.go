package models

import (
	"errors"
	"regexp"
	"time"
)

type Student struct {
	User                                   // Composition
	AdmissionNumber    string              `gorm:"uniqueIndex" json:"admission_number"`
	Campus             string              `json:"campus"`
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

func (u *Student) Validate() error {
	// Set default campus for Student
	if u.Campus == "" {
		u.Campus = "athi"
	}
	// user validations
	if err := u.User.Validate(); err != nil {
		return err
	}
	// Username
	if len(u.AdmissionNumber) < 7 && !VerifyAdmno(u.AdmissionNumber) {
		return errors.New("Please provide a valid admission number in the 00-0000 format")
	}
	if u.Campus != "athi" && u.Campus != "nairobi" {
		return errors.New("Please specify a valid campus athi or nairobi")
	}

	return nil
}
