package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Realm struct {
	BaseModel
	Name        string `gorm:"not null;unique" validate:"required"`
	Description string `gorm:"not null" validate:"required"`
}

func (r *Realm) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
