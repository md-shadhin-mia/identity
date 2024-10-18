package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdministrativeUser struct {
	BaseModel
	UserId uuid.UUID
	User   User `gorm:"foreignKey:UserId"`
}

func (u *AdministrativeUser) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
