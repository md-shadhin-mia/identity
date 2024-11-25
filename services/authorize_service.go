package services

import (
	"errors"
	"indentity/models"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthorizeService struct {
	db *gorm.DB
}

func NewAuthorizeService(db *gorm.DB) *AuthorizeService {
	return &AuthorizeService{db: db}
}

func (a *AuthorizeService) Authorize(client_id string, redirect_uri string, scope string, state string) {

}
func (a *AuthorizeService) ClientCredentialsGrant(client_id string, client_secret string, scope string) (string, error) {
	var client models.Client
	if err := a.db.Where("id = ? AND secret = ?", client_id, client_secret).First(&client).Error; err != nil {
		return "", err
	}
	token, err := a.generateAccessToken(client)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthorizeService) PasswordGrant(client_id string, client_secret string, username string, password string, scope string) (string, error) {
	var user models.User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}
	if !user.CheckPasswordHash(password) {
		return "", errors.New("invalid password")
	}

	// Todo: check user exists admin role check
	if !user.HasRole("admin") {
		return "", errors.New("user is not an admin")
	}

	// Todo: load user permissions
	permissions, err := a.loadUserPermissions(user.ID)
	if err != nil {
		return "", err
	}

	// Todo: generate token
	token, err := a.generateToken(user, permissions, scope)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthorizeService) loadUserPermissions(userID uint) ([]string, error) {
	// implement loading user permissions from database or cache
	// for example:
	var permissions []models.Permission
	if err := a.db.Where("user_id = ?", userID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	var permissionStrings []string
	for _, permission := range permissions {
		permissionStrings = append(permissionStrings, permission.Name)
	}
	return permissionStrings, nil
}

func (a *AuthorizeService) generateToken(user models.User, permissions []string, scope string) (string, error) {
	// implement generating token using a library like jwt-go
	// for example:
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         user.ID,
		"permissions": permissions,
		"scope":       scope,
	})
	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *AuthorizeService) generateAccessToken(client models.Client) (string, error) {

	return "", nil
}

func (a *AuthorizeService) generateRefreshToken(client models.Client) (string, error) {
	return "", nil
}
