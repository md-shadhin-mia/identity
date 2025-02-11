// client_controller.go
package controllers

import (
	"indentity/models"
	"indentity/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClientInput struct {
	Name   string `json:"name" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

type ClientController struct {
	cs   *services.ClientService
	rbac *services.RBACService
}

func NewClientController(cs *services.ClientService, rbac *services.RBACService) *ClientController {
	return &ClientController{cs: cs, rbac: rbac}
}

func (cc *ClientController) CreateClient(c *gin.Context) {
	var clientInput ClientInput
	err := c.BindJSON(&clientInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &models.Client{
		Name:   clientInput.Name,
		Secret: clientInput.Secret,
	}

	err = cc.cs.CreateClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func (cc *ClientController) GetClientByID(c *gin.Context) {
	id := c.Param("id")

	client, err := cc.cs.GetClientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (cc *ClientController) GetClientByName(c *gin.Context) {
	name := c.Param("name")

	client, err := cc.cs.GetClientByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (cc *ClientController) UpdateClient(c *gin.Context) {
	id := c.Param("id")

	var clientInput ClientInput
	err := c.BindJSON(&clientInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := cc.cs.GetClientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	client.Name = clientInput.Name
	client.Secret = clientInput.Secret

	err = cc.cs.UpdateClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (cc *ClientController) DeleteClient(c *gin.Context) {
	id := c.Param("id")

	err := cc.cs.DeleteClient(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

type AssignRoleInput struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
	RoleID uuid.UUID `json:"role_id" binding:"required"`
}

func (cc *ClientController) AssignRole(c *gin.Context) {
	clientID := c.Param("client_id")
	var assignRoleInput AssignRoleInput

	err := c.BindJSON(&assignRoleInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientIDUuid, err := uuid.Parse(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	err = cc.rbac.AssignRole(clientIDUuid, assignRoleInput.UserID, assignRoleInput.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role assigned successfully"})
}

type RevokeRoleInput struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
	RoleID uuid.UUID `json:"role_id" binding:"required"`
}

func (cc *ClientController) RevokeRole(c *gin.Context) {
	clientID := c.Param("client_id")
	var revokeRoleInput RevokeRoleInput

	err := c.BindJSON(&revokeRoleInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientIDUuid, err := uuid.Parse(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	err = cc.rbac.RevokeRole(clientIDUuid, revokeRoleInput.UserID, revokeRoleInput.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role revoked successfully"})
}

type CreateRoleInput struct {
	Name        string   `json:"name" binding:"required"`
	Permissions []string `json:"permissions" binding:"required"`
}

func (cc *ClientController) CreateClientRole(c *gin.Context) {
	clientID := c.Param("client_id")
	var createRoleInput CreateRoleInput

	err := c.BindJSON(&createRoleInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientIDUuid, err := uuid.Parse(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	err = cc.rbac.CreateClientRole(clientIDUuid, createRoleInput.Name, createRoleInput.Permissions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role created successfully"})
}
