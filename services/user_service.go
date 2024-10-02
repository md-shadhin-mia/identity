package services

import (
	"indentity/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (us *UserService) CreateUser(user *models.User) error {
	user.SetPassword(user.Password)
	return us.db.Create(user).Error
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := us.db.First(&user, "id = ?", uuid.MustParse(id)).Error
	return &user, err
}

// func (us *UserService) AddRoleToUser(userID string, roleID uint) error {
// 	userRole := &models.UserRole{RoleID: roleID}
// 	return us.db.Create(userRole).Error
// }

func (us *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := us.db.First(&user, "username = ?", username).Error
	return &user, err
}
