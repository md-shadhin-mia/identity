package models

import (
	"indentity/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Name   string    `gorm:"uniqueIndex;not null"`
	Secret string    `gorm:"not null"`
}

func (c *Client) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	c.Secret = utils.GenerateSecret()
	return nil
}

func (c *Client) SetSecret(secret string) {
	c.Secret = secret
}

func (c *Client) CheckSecret(secret string) bool {
	return c.Secret == secret
}
