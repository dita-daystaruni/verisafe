package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	User        uuid.UUID `gorm:"type:uuid" json:"user"`
	TokenString string    `gorm:"unique;not null" json:"token_string"`
	DateCreated time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date_created"`
}

// Deletes all the user's preexisting tokens and
// append a new token onto the database.
func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.DateCreated = time.Now()

	// Delete all existing tokens for the same user
	if err := tx.Unscoped().Where("tokens.user = ?", t.User).Delete(&Token{}).Error; err != nil {
		return err
	}

	return nil
}
