package models

import (
	"indentity/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	BaseModel
	Name      string       `gorm:"uniqueIndex;not null"`
	Secret    string       `gorm:"not null"`
	Roles     []ClientRole `gorm:"foreignKey:ClientID"`
	UserRoles []UserRole   `gorm:"foreignKey:ClientID"`
}

func (c *Client) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	c.Secret = utils.GenerateSecret()
	return nil
}

func (c *Client) SetSecret(secret string) {
	c.Secret = secret
}

func (c *Client) CheckSecret(secret string) bool {
	return c.Secret == secret
}
