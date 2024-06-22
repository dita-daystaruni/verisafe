package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RewardTransaction struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	StudentID uuid.UUID `gorm:"type:uuid" json:"student_id"`
	Points    int       `json:"points"`
	Reason    string    `json:"reason"`
	AwardedAt time.Time `json:"awarded_at"`
}

// hooks
func (rt *RewardTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if rt.Points > 10 {
		return errors.New("A transaction can only have a maximum of 10 points per request")
	}

	if len(rt.Reason) < 5 {
		return errors.New("Please specify a valid reason for allocating points")
	}

	return nil
}

func (rt *RewardTransaction) BeforeUpdate(tx *gorm.DB) (err error) {
	if rt.Points > 10 {
		return errors.New("A transaction can only have a maximum of 50 points per request")
	}

	return nil
}
