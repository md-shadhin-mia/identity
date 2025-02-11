package middleware

import (
	"indentity/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RBACMiddleware checks if the user has the required permission for the client
func RBACMiddleware(rbac *services.RBACService, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context (assuming it's already set by an authentication middleware)
		userID, err := getUserIDFromContext(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID not found"})
			return
		}

		// Get the client ID from the request parameters
		clientID, err := getClientIDFromParams(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Invalid Client ID"})
			return
		}

		// Check if the user has the required permission for the client
		if !rbac.CheckPermission(clientID, userID, permission) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient permissions"})
			return
		}

		// If the user has the required permission, continue to the next handler
		c.Next()
	}
}

// Helper function to extract the user ID from the context
func getUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	userIDString := c.GetHeader("User-ID") // Assuming the user ID is passed in the "User-ID" header
	if userIDString == "" {
		return uuid.Nil, gin.Error{Err: nil, Type: gin.ErrorTypePrivate, Meta: "User ID not found in header"}
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

// Helper function to extract the client ID from the request parameters
func getClientIDFromParams(c *gin.Context) (uuid.UUID, error) {
	clientIDString := c.Param("client_id") // Assuming the client ID is passed as a parameter named "client_id"
	if clientIDString == "" {
		return uuid.Nil, gin.Error{Err: nil, Type: gin.ErrorTypePrivate, Meta: "Client ID not found in parameters"}
	}

	clientID, err := uuid.Parse(clientIDString)
	if err != nil {
		return uuid.Nil, err
	}

	return clientID, nil
}

// ScopeMiddleware checks if the request has a specific scope
func ScopeMiddleware(scope string) gin.HandlerFunc {

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
		// token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// 	}
		// 	return secret, nil
		// })
		// if err != nil {
		// 	log.Println(err)
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token sfa"})
		// 	c.Abort()
		// 	return
		// }

		// // Get the scopes from the token claims
		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok || !token.Valid {
		// 	log.Println("claims are not valid")
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token! claims are not valid"})
		// 	c.Abort()
		// 	return
		// }
		// scopes, ok := claims["scopes"].([]string)
		// if !ok {
		// 	log.Println("scopes not founded in token")
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token! scopes not founded in token"})
		// 	c.Abort()
		// 	return
		// }

		// // Check if the scope is present
		// if !contains(scopes, scope) {
		// 	log.Println("scope is not valid")
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden ! scope is not valid"})
		// 	c.Abort()
		// 	return
		// }

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
