package services

import (
	"errors"
	"indentity/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthenticationService struct {
	db *gorm.DB
}

var secret = []byte("secret")
var AllScopes = []string{"reamls:read", "reamls:write"}

func NewAuthenticationService(db *gorm.DB) *AuthenticationService {
	secret = []byte(os.Getenv("JWT_SECRET"))
	return &AuthenticationService{db: db}
}

// AuthenticateUser authenticates a user based on their credentials.
func (as *AuthenticationService) AuthenticateUser(username string, password string) (map[string]interface{}, error) {
	user := models.User{}
	err := as.db.First(&user, models.User{Username: username}).Error
	if err != nil {
		return nil, err
	}

	// Check if the password is valid
	if !user.CheckPasswordHash(password) {
		return nil, errors.New("invalid password")
	}

	// Generate tokens
	accessToken, err := as.generateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := as.generateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	idToken, err := as.generateIDToken(user.ID)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"access_token":       accessToken,
		"expires_in":         3600, // 1 hour
		"refresh_expires_in": 1800, // 30 minutes
		"refresh_token":      refreshToken,
		"token_type":         "Bearer",
		"id_token":           idToken,
	}

	return response, nil
}

// generateAccessToken generates an access token for the given user ID.
func (as *AuthenticationService) generateAccessToken(user models.User) (string, error) {
	// Generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = time.Now().Unix()
	claims["sub"] = user.ID.String()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	if user.Username == "admin" {
		claims["scopes"] = AllScopes
	}
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// generateRefreshToken generates a refresh token for the given user ID.
func (as *AuthenticationService) generateRefreshToken(userID uuid.UUID) (string, error) {
	// Generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID.String()
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// generateIDToken generates an ID token for the given user ID.
func (as *AuthenticationService) generateIDToken(userID uuid.UUID) (string, error) {
	// Generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID.String()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
