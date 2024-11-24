// client_controller.go
package controllers

import (
	"indentity/core"
	"indentity/models"
	"indentity/services"

	"github.com/gin-gonic/gin"
)

type ClientInput struct {
	Name   string `json:"name" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

type ClientController struct {
	cs *services.ClientService
}

func NewClientController(cs *services.ClientService) *ClientController {
	return &ClientController{cs: cs}
}

func (cc *ClientController) CreateClient(c *core.Context) {
	var clientInput ClientInput
	err := c.BindJSON(&clientInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	client := &models.Client{
		Name:   clientInput.Name,
		Secret: clientInput.Secret,
	}

	err = cc.cs.CreateClient(client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, client)
}

func (cc *ClientController) GetClientByID(c *core.Context) {
	id := c.Param("id")

	client, err := cc.cs.GetClientByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, client)
}

func (cc *ClientController) GetClientByName(c *core.Context) {
	name := c.Param("name")

	client, err := cc.cs.GetClientByName(name)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, client)
}

func (cc *ClientController) UpdateClient(c *core.Context) {
	id := c.Param("id")

	var clientInput ClientInput
	err := c.BindJSON(&clientInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	client, err := cc.cs.GetClientByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	client.Name = clientInput.Name
	client.Secret = clientInput.Secret

	err = cc.cs.UpdateClient(client)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, client)
}

func (cc *ClientController) DeleteClient(c *core.Context) {
	id := c.Param("id")

	err := cc.cs.DeleteClient(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Client deleted successfully"})
}
