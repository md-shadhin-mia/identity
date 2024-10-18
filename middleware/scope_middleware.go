package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

// ScopeMiddleware checks if the request has a specific scope
func ScopeMiddleware(scope string) gin.HandlerFunc {

	secret := []byte(os.Getenv("JWT_SECRET"))
	return func(c *gin.Context) {
		// Get the Bearer token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract the token from the Bearer header
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		// log.Print("tokenString: ", tokenString)
		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token sfa"})
			c.Abort()
			return
		}

		// Get the scopes from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("claims are not valid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token! claims are not valid"})
			c.Abort()
			return
		}
		scopes, ok := claims["scopes"].([]string)
		if !ok {
			log.Println("scopes not founded in token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token! scopes not founded in token"})
			c.Abort()
			return
		}

		// Check if the scope is present
		if !contains(scopes, scope) {
			log.Println("scope is not valid")
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden ! scope is not valid"})
			c.Abort()
			return
		}

		// If the scope is present, continue to the next handler
		c.Next()
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
