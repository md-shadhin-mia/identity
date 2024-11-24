package middleware

import (
	"fmt"
	"indentity/core"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

// ScopeMiddleware checks if the request has a specific scope
func ScopeMiddleware(scope string) core.HandlerFunc {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return func(c *core.Context) {
		// Get the Bearer token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Extract the token from the Bearer header
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if err != nil {
			log.Printf("Token validation error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
			c.Abort()
			return
		}

		// Get the claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("Failed to parse token claims or token is invalid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			c.Abort()
			return
		}

		// Get scopes from claims and ensure it's an array of strings
		scopesInterface, exists := claims["scopes"]
		if !exists {
			log.Printf("Token claims: %+v scopes not found", claims)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: scopes not found in token"})
			c.Abort()
			return
		}

		// Convert interface{} to []string
		var scopes []string
		if scopesArray, ok := scopesInterface.([]interface{}); ok {
			scopes = make([]string, len(scopesArray))
			for i, s := range scopesArray {
				if str, ok := s.(string); ok {
					scopes[i] = str
				} else {
					log.Printf("Invalid scope type at index %d: %T", i, s)
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: malformed scopes"})
					c.Abort()
					return
				}
			}
		} else {
			log.Printf("Invalid scopes type: %T", scopesInterface)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: malformed scopes"})
			c.Abort()
			return
		}

		// Check if the scope is present
		if !contains(scopes, scope) {
			log.Printf("Required scope '%s' not found in token scopes %v", scope, scopes)
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Forbidden: missing required scope '%s'", scope)})
			c.Abort()
			return
		}

		// If the scope is present, continue to the next handler
		c.Next()
	}
}

// contains checks if a string is present in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
