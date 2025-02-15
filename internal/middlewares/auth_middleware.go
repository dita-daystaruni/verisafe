package middlewares

import (
	"errors"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

// Claims structure for JWT
type VerisafeClaims struct {
	User        repository.GetUserByIDRow `json:"user"`
	Roles       []string                  `json:"roles"`
	Permissions []string                  `json:"permissions"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new token for a given user ID.
func GenerateJWT(
	user repository.GetUserByIDRow,
	roles []repository.ListRolesForUserRow,
	permissions []string,
	cfg configs.Config) (string, error) {

	roleNames := []string{}
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}

	claims :=
		&VerisafeClaims{
			User:        user,
			Roles:       roleNames,
			Permissions: permissions,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWTConfig.ExpireDelta))),
				Issuer:    "verisafe",
			},
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTConfig.ApiSecret))
}

// ValidateJWT parses and validates the JWT token and checks expiration.
func ValidateJWT(tokenString string, secret string) (*VerisafeClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &VerisafeClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	claims, ok := token.Claims.(*VerisafeClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token you have. Create a valid one you must!")
	}

	// Check if the token is expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Your token expired it is. Refresh it you must")
	}

	return claims, nil
}
