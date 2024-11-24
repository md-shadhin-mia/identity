package controllers

// controllers/reaml_controller.go

import (
	"indentity/core"
	"indentity/models"
	"indentity/services"
	"log"

	"github.com/gin-gonic/gin"
)

type ReamlController struct {
	rs *services.RealmService
}

func NewReamlController(rs *services.RealmService) *ReamlController {
	return &ReamlController{rs: rs}
}

type ReamlInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (rc *ReamlController) GetAllReamls(c *core.Context) {
	reamls, err := rc.rs.GetAllReamls()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, reamls)
}
func (rc *ReamlController) CreateReaml(c *core.Context) {
	var reamlInput models.Realm
	err := c.BindValidateJson(&reamlInput)
	if err != nil {
		// c.JSON(400, gin.H{"error": err})
		log.Println(err)
		return
	}
	reaml, err := rc.rs.CreateRealm(&reamlInput)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, reaml)
}

func (rc *ReamlController) GetReamlByID(c *core.Context) {
	id := c.Param("id")
	reaml, err := rc.rs.GetRealmByName((id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, reaml)
}
func (rc *ReamlController) UpdateReaml(c *core.Context) {
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

func (rc *ReamlController) DeleteReaml(c *core.Context) {
	id := c.Param("id")

	err := rc.rs.DeleteRealm(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Reaml deleted successfully"})
}
