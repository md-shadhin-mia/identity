package controllers

// controllers/reaml_controller.go

import (
	"indentity/models"
	"indentity/services"

	"github.com/gin-gonic/gin"
)

type ReamlController struct {
	rs *services.RealmService
}

func NewReamlController(rs *services.RealmService) *ReamlController {
	return &ReamlController{rs: rs}
}

type ReamlInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (rc *ReamlController) CreateReaml(c *gin.Context) {
	var reamlInput ReamlInput
	err := c.BindJSON(&reamlInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = rc.rs.CreateRealm(&models.Realm{Name: reamlInput.Name, Description: reamlInput.Description})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, reamlInput)
}

func (rc *ReamlController) GetReamlByID(c *gin.Context) {
	id := c.Param("id")
	reaml, err := rc.rs.GetRealmByName((id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, reaml)
}
func (rc *ReamlController) UpdateReaml(c *gin.Context) {
	id := c.Param("id")

	var reamlInput ReamlInput
	err := c.BindJSON(&reamlInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	reaml, err := rc.rs.GetRealmByName(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	reaml.Name = reamlInput.Name
	reaml.Description = reamlInput.Description

	err = rc.rs.UpdateRealm(reaml)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, reaml)
}

func (rc *ReamlController) DeleteReaml(c *gin.Context) {
	id := c.Param("id")

	err := rc.rs.DeleteRealm(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Reaml deleted successfully"})
}
