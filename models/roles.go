package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRole struct {
	BaseModel
	ClientID    uuid.UUID
	Name        string       `gorm:"not null"`
	Permissions []Permission `gorm:"foreignKey:RoleID"`
}

type Permission struct {
	BaseModel
	RoleID uuid.UUID
	Name   string `gorm:"not null"`
}

type UserRole struct {
	BaseModel
	ClientID uuid.UUID
	UserID   uuid.UUID
	RoleID   uuid.UUID
}

func (r *ClientRole) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (ur *UserRole) BeforeCreate(tx *gorm.DB) error {
	if ur.ID == uuid.Nil {
		ur.ID = uuid.New()
	}
	return nil
}
