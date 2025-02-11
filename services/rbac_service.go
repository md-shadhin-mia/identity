package services

import (
	"indentity/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RBACService struct {
	db *gorm.DB
}

func NewRBACService(db *gorm.DB) *RBACService {
	return &RBACService{db: db}
}

func (r *RBACService) AssignRole(clientID uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error {
	userRole := models.UserRole{
		ClientID: clientID,
		UserID:   userID,
		RoleID:   roleID,
	}

	result := r.db.Create(&userRole)
	return result.Error
}

func (r *RBACService) RevokeRole(clientID uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error {
	result := r.db.Where("client_id = ? AND user_id = ? AND role_id = ?", clientID, userID, roleID).Delete(&models.UserRole{})
	return result.Error
}

func (r *RBACService) CheckPermission(clientID uuid.UUID, userID uuid.UUID, permissionName string) bool {
	var userRole models.UserRole
	result := r.db.Where("client_id = ? AND user_id = ?", clientID, userID).First(&userRole)
	if result.Error != nil {
		return false
	}

	var role models.ClientRole
	result = r.db.Preload("Permissions").First(&role, userRole.RoleID)
	if result.Error != nil {
		return false
	}

	for _, permission := range role.Permissions {
		if permission.Name == permissionName {
			return true
		}
	}

	return false
}

func (r *RBACService) CreateClientRole(clientID uuid.UUID, roleName string, permissions []string) error {
	clientRole := models.ClientRole{
		ClientID: clientID,
		Name:     roleName,
	}

	result := r.db.Create(&clientRole)
	if result.Error != nil {
		return result.Error
	}

	for _, permissionName := range permissions {
		permission := models.Permission{
			RoleID: clientRole.ID,
			Name:   permissionName,
		}
		result = r.db.Create(&permission)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
