package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dita-daystaruni/verisafe/internal/configs"
	"github.com/dita-daystaruni/verisafe/internal/repository"
	"github.com/gin-gonic/gin"
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
	if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Your token expired it is. Refresh it you must")
	}

	return claims, nil
}

// PermissionMiddleware checks if the user has the required permissions to access a resource
func PermissionMiddleware(requiredPermissions []string, cfg *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Split the Bearer token from the header
		tokenString := strings.Split(authHeader, " ")[1]

		// Validate and parse the JWT token
		claims, err := ValidateJWT(tokenString, cfg.JWTConfig.ApiSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Check if the user has the required permissions
		hasPermission := false
		for _, userPermission := range claims.Permissions {
			for _, requiredPermission := range requiredPermissions {
				if userPermission == requiredPermission {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden,
				gin.H{"error": "You do not have the required permissions to access this resource"},
			)
			c.Abort()
			return
		}

		// User is authenticated and authorized, proceed with the request
		c.Next()
	}
}

func RoleAndPermissionMiddleware(requiredRoles []string,
	requiredPermissions []string,
	cfg configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header and validate the JWT token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := ValidateJWT(tokenString, cfg.JWTConfig.ApiSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Check if the user has the required roles
		hasRole := false
		for _, userRole := range claims.Roles {
			for _, requiredRole := range requiredRoles {
				if userRole == requiredRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		// If no required roles, return forbidden
		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required roles to access this resource"})
			c.Abort()
			return
		}

		// Check if the user has the required permissions
		hasPermission := false
		for _, userPermission := range claims.Permissions {
			for _, requiredPermission := range requiredPermissions {
				if userPermission == requiredPermission {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required permissions to access this resource"})
			c.Abort()
			return
		}

		// Proceed with the request if roles and permissions are valid
		c.Next()
	}
}
