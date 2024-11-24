package services

import "gorm.io/gorm"

type AuthorizeService struct {
	db *gorm.DB
}

func NewAuthorizeService(db *gorm.DB) *AuthorizeService {
	return &AuthorizeService{db: db}
}
