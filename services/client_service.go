// client_service.go
package services

import (
	"indentity/models"
	"indentity/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientService struct {
	db *gorm.DB
}

func NewClientService(db *gorm.DB) *ClientService {
	return &ClientService{db: db}
}

func (cs *ClientService) CreateClient(client *models.Client) error {
	client.SetSecret(utils.GenerateSecret())
	return cs.db.Create(client).Error
}

func (cs *ClientService) GetClientByID(id string) (*models.Client, error) {
	var client models.Client
	err := cs.db.First(&client, "id = ?", uuid.MustParse(id)).Error
	return &client, err
}

func (cs *ClientService) GetClientByName(name string) (*models.Client, error) {
	var client models.Client
	err := cs.db.First(&client, "name = ?", name).Error
	return &client, err
}

func (cs *ClientService) UpdateClient(client *models.Client) error {
	return cs.db.Save(client).Error
}

func (cs *ClientService) DeleteClient(id string) error {
	var client models.Client
	client.ID = uuid.MustParse(id)
	return cs.db.Delete(&client).Error
}
