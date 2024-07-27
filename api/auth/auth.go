package auth

import (
	"errors"
	"time"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/models"
	"github.com/dita-daystaruni/verisafe/models/db"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// Generates a jwt token that can be used to authenticate
// a user
func GenerateToken(user *models.User, isAdmin bool, cfg *configs.Config, con *gorm.DB) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"username":   user.Username,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"is_admin":   isAdmin,
		"exp":        time.Now().Add(time.Duration(cfg.JWTConfig.ExpireDelta) * time.Minute).Unix(),
	})

	// Sign with the API secret
	tokenString, err := claims.SignedString([]byte(cfg.JWTConfig.ApiSecret))
	if err != nil {
		return "", err
	}

	ts := db.TokenStore{DB: con}

	token, err := ts.RegisterToken(models.Token{User: user.ID, TokenString: tokenString})
	if err != nil {
		return "", err
	}

	return token.TokenString, nil
}

// Verifies a user's token
func VerifyToken(tokenString string, cfg *configs.Config) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTConfig.ApiSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("You provided an invalid token please relogin to continue")
	}

	return token, nil
}
