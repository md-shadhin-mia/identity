package controllers

import (
	"indentity/core"
	"indentity/services"

	"github.com/gin-gonic/gin"
)

type AuthorizeController struct {
	service *services.AuthorizeService
}

func NewAuthorizeController(service *services.AuthorizeService) *AuthorizeController {
	return &AuthorizeController{service: service}
}

func (a *AuthorizeController) Authorize(c *core.Context) {
	response_type := c.Query("response_type")
	client_id := c.Query("client_id")
	redirect_uri := c.Query("redirect_uri")
	scope := c.Query("scope")
	state := c.Query("state")
	if response_type == "code" {
		//Todo: a login page
		c.HTML(200, "login_page.html", gin.H{"client_id": client_id, "redirect_uri": redirect_uri, "scope": scope, "state": state})
	} else if response_type == "token" {
		//Todo: a login page
		c.HTML(200, "login_page", gin.H{"client_id": client_id, "redirect_uri": redirect_uri, "scope": scope, "state": state})
	}
}

type TokenInput struct {
	Grant_type    string  `json:"grant_type" validate:"required"`
	Client_id     string  `json:"client_id" validate:"required"`
	Client_secret string  `json:"client_secret" validate:"required"`
	Code          *string `json:"code" validate:"omitempty"`
	Username      *string `json:"username" validate:"omitempty"`
	Password      *string `json:"password" validate:"omitempty"`
	Scope         *string `json:"scope" validate:"omitempty"`
}

func (a *AuthorizeController) TokenAuthorize(c *core.Context) {
	tokenInput := TokenInput{}
	err := c.BindValidateJson(&tokenInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if tokenInput.Grant_type == "password" {
		if tokenInput.Username == nil || tokenInput.Password == nil {
			c.JSON(400, gin.H{"error": "username or password is required"})
			return
		}
		r, err := a.service.PasswordGrant(tokenInput.Client_id, tokenInput.Client_secret, *tokenInput.Username, *tokenInput.Password, *tokenInput.Scope)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"token": r})
	}
}
