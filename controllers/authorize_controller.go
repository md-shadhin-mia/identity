package controllers

import (
	"indentity/core"
	"indentity/services"

	"github.com/gin-gonic/gin"
)

type AuthorizeController struct {
	service services.AuthorizeService
}

func NewAuthorizeController(service services.AuthorizeService) *AuthorizeController {
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
		c.HTML(200, "login.tmpl", gin.H{"client_id": client_id, "redirect_uri": redirect_uri, "scope": scope, "state": state})
	} else if response_type == "token" {
		//Todo: a login page
		c.HTML(200, "login.tmpl", gin.H{"client_id": client_id, "redirect_uri": redirect_uri, "scope": scope, "state": state})
	} else {
		c.JSON(400, gin.H{"error": "unsupported_response_type", "error_description": "The authorization server does not support obtaining an access token using this method."})
	}
}
