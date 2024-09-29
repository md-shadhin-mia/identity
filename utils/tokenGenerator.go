package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signingKey = []byte("your-256-bit-secret")

// Function to generate ID token
func GenerateIDToken() (string, error) {
	claims := jwt.MapClaims{
		"iss":       "https://auth.example.com",
		"sub":       "1234567890",
		"aud":       "client_id_example",
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
		"iat":       time.Now().Unix(),
		"auth_time": time.Now().Unix(),
		"email":     "user@example.com",
		"name":      "John Doe",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// Function to generate Access token
func GenerateAccessToken() (string, error) {
	claims := jwt.MapClaims{
		"iss":       "https://auth.example.com",
		"sub":       "1234567890",
		"aud":       "https://api.example.com",
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
		"scope":     "read write",
		"client_id": "client_id_example",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// Function to generate Refresh token (opaque token)
func GenerateRefreshToken() string {
	// Refresh tokens are usually opaque and generated using random strings
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func GenerateSecret() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
