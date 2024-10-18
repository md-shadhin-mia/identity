package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	// Token expiration time in seconds
	tokenExpiration = 3600
)

// AuthenticationService provides methods for authentication
type AuthenticationService struct {
	secretKey string
}

// NewAuthenticationService returns a new AuthenticationService instance
func NewAuthenticationService(secretKey string) *AuthenticationService {
	return &AuthenticationService{secretKey: secretKey}
}

// AuthorizationCodeGrant handles the authorization code grant flow
func (a *AuthenticationService) AuthorizationCodeGrant(w http.ResponseWriter, r *http.Request) {
	// Get the authorization code from the request
	code := r.FormValue("code")

	// Validate the authorization code
	if code == "" {
		http.Error(w, "Invalid authorization code", http.StatusBadRequest)
		return
	}

	// Generate an access token
	token := a.generateAccessToken()

	// Return the access token
	w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, token)))
}

// ImplicitGrant handles the implicit grant flow
func (a *AuthenticationService) ImplicitGrant(w http.ResponseWriter, r *http.Request) {
	// Get the client ID and redirect URI from the request
	clientID := r.FormValue("client_id")
	redirectURI := r.FormValue("redirect_uri")

	// Validate the client ID and redirect URI
	if clientID == "" || redirectURI == "" {
		http.Error(w, "Invalid client ID or redirect URI", http.StatusBadRequest)
		return
	}

	// Generate an access token
	token := a.generateAccessToken()

	// Redirect the user to the redirect URI with the access token
	http.Redirect(w, r, fmt.Sprintf("%s#access_token=%s", redirectURI, token), http.StatusFound)
}

// ClientCredentialsGrant handles the client credentials grant flow
func (a *AuthenticationService) ClientCredentialsGrant(w http.ResponseWriter, r *http.Request) {
	// Get the client ID and client secret from the request
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	// Validate the client ID and client secret
	if clientID == "" || clientSecret == "" {
		http.Error(w, "Invalid client ID or client secret", http.StatusBadRequest)
		return
	}

	// Generate an access token
	token := a.generateAccessToken()

	// Return the access token
	w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, token)))
}

// ResourceOwnerPasswordGrant handles the resource owner password grant flow
func (a *AuthenticationService) ResourceOwnerPasswordGrant(w http.ResponseWriter, r *http.Request) {
	// Get the username and password from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validate the username and password
	if username == "" || password == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	// Generate an access token
	token := a.generateAccessToken()

	// Return the access token
	w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, token)))
}

// RefreshTokenGrant handles the refresh token grant flow
func (a *AuthenticationService) RefreshTokenGrant(w http.ResponseWriter, r *http.Request) {
	// Get the refresh token from the request
	refreshToken := r.FormValue("refresh_token")

	// Validate the refresh token
	if refreshToken == "" {
		http.Error(w, "Invalid refresh token", http.StatusBadRequest)
		return
	}

	// Generate a new access token
	token := a.generateAccessToken()

	// Return the new access token
	w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, token)))
}

// generateAccessToken generates a new access token
func (a *AuthenticationService) generateAccessToken() string {
	// Generate a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * tokenExpiration).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}
