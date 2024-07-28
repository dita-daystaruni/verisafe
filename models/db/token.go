package db

import (
	"github.com/dita-daystaruni/verisafe/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenStore struct {
	DB *gorm.DB
}

// Saves a user's token onto the databse
func (ts *TokenStore) RegisterToken(token models.Token) (*models.Token, error) {
	if err := ts.DB.Save(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

// Retrieves a token specified by [tokenString] from the databse
func (ts *TokenStore) RetrieveToken(tokenString string) (*models.Token, error) {
	var token models.Token

	err := ts.DB.Where("token_string = ?", tokenString).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// Deletes a token specified by tokenString
func (ts *TokenStore) DeleteToken(tokenString string) error {

	err := ts.DB.Unscoped().Where("token_string = ?", tokenString).Delete(&models.Token{}).Error
	if err != nil {
		return err
	}

	return nil
}

// Killswitch to delete all tokens and refresh causing users to
// relog in
func (ts *TokenStore) DeleteAllTokens(tokenID uuid.UUID) error {
	err := ts.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Token{}).Error
	if err != nil {
		return err
	}

	return nil
}
