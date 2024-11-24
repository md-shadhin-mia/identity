package controllers

import (
	"indentity/core"
	"indentity/models"
	"indentity/services"
	"indentity/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{us: us}
}

func (uc *UserController) GetAllUsers(c *core.Context) {
	users, err := uc.us.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uc *UserController) CreateUser(c *core.Context) {
	var user UserInput
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = uc.us.CreateUser(&models.User{Username: user.Username, Email: user.Email, Password: user.Password})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}

func (uc *UserController) GetUserByID(c *core.Context) {
	id := c.Param("id")

	user, err := uc.us.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// func (uc *UserController) AddRoleToUser(c *core.Context) {
// 	userID := c.Param("user_id")
// 	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err = uc.us.AddRoleToUser(userID, uint(roleID))
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "Role added to user successfully"})
// }

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uc *UserController) SignIn(c *core.Context) {
	var user LoginRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userFromDB, err := uc.us.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if !userFromDB.CheckPasswordHash(user.Password) {

		c.JSON(401, gin.H{"error": "invalid password"})
		return
	}
	token, err := utils.GenerateToken(userFromDB.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (uc *UserController) AuthenticateUser(c *core.Context) {
	var user LoginRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	r, err := uc.us.AuthenticateUser(user.Username, user.Password)

	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(200, gin.H{"token": r})
}
