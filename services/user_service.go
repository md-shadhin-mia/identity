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

/* get all users */
func (us *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := us.db.Find(&users).Error
	return users, err
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

func (us *UserService) AuthenticateUser(username string, password string) (map[string]interface{}, error) {
	AuthenticationService := NewAuthenticationService(us.db)
	return AuthenticationService.AuthenticateUser(username, password)
}
