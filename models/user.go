package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(191);not null" json:"-"`

	RealmID uuid.UUID `gorm:"type:binary(16)"`
	Realm   *Realm    `gorm:"foreignkey:RealmID;association_foreignkey:ID"`
}

type Role struct {
	BaseModel
	Name string `gorm:"uniqueIndex;not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
