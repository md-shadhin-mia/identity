// services/user_service.go

package services

import (
	"indentity/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RealmService struct {
	db *gorm.DB
}

func NewRealmService(db *gorm.DB) *RealmService {
	return &RealmService{db: db}
}

func (rs *RealmService) CreateRealm(realm *models.Realm) error {
	return rs.db.Create(realm).Error
}

func (rs *RealmService) GetRealmByName(name string) (*models.Realm, error) {
	var realm models.Realm
	err := rs.db.First(&realm, "name = ?", name).Error
	return &realm, err
}

func (rs *RealmService) UpdateRealm(realm *models.Realm) error {
	return rs.db.Save(realm).Error
}

func (rs *RealmService) DeleteRealm(id string) error {
	var realm models.Realm
	realm.ID = uuid.MustParse(id)
	return rs.db.Delete(&realm).Error
}
