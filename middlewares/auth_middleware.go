package middlewares

import (
	"errors"
	"time"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

// Replace this with a secure random secret key in production.
var jwtSecret = []byte("your-secret-key")

// Claims structure for JWT
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new token for a given user ID.
func GenerateJWT(user repository.User, cfg configs.Config) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":    user.ID.String(),
		"username":   user.Username,
		"email":      user.Email,
		"firstname":  user.Firstname,
		"expires_at": time.Now().Add(time.Hour * time.Duration(cfg.JWTConfig.ExpireDelta)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTConfig.ApiSecret))
}

// ValidateJWT parses and validates the JWT token and checks expiration.
func ValidateJWT(tokenString string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with the expected method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and validate the claims
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token you have create a valid one you must!")
	}

	// Check if the token is expired
	if exp, ok := (*claims)["expires_at"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errors.New("Your token expired it is. Refresh it you must")
		}
	} else {
		return nil, errors.New("Your token wrong it is. Refresh it you must")
	}

	return claims, nil
}
