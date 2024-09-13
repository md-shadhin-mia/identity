package main

import (
	"indentity/controllers"
	"indentity/initilizer"
	"indentity/services"
	"indentity/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizer.LoadEnv()
	initilizer.DBConnect()
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	var v1 = r.Group("/api/v1")
	userService := services.NewUserService(initilizer.DB)

	userController := controllers.NewUserController(userService)
	v1.POST("/users", userController.CreateUser)

	getUser := v1.Group("/users")
	getUser.Use(utils.TokenMiddleware())
	getUser.GET("/:id", userController.GetUserByID)

	v1.POST("/users/:user_id/roles/:role_id", userController.AddRoleToUser)

	v1.POST("/signin", userController.SignIn)

	// // utils.Resources(*v1, "users", controllers.NewUserController(initilizer.DB))
	// utils.Resources(*v1, "demos", controllers.NewDemoController(initilizer.DB))
	// // utils.Resources(*v1, "roles", &controllers.RoleController{DB: initilizer.DB})
	// // v1.POST("/users/:id/add_roles", controllers.NewUserController(initilizer.DB).AddRoles)

	// // r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
