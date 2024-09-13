package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uuid.UUID `gorm:"not null"`
	RoleID uint      `gorm:"index;not null"`
	User   User      `gorm:"foreignkey:UserID"`
	Role   Role      `gorm:"foreignkey:RoleID"`
}
