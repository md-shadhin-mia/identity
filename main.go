package main

import (
	"indentity/controllers"
	"indentity/initilizer"
	"indentity/middleware"
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

	reamlService := services.NewRealmService(initilizer.DB)
	reamlController := controllers.NewReamlController(reamlService)
	v1.POST("/reamls", middleware.ScopeMiddleware("realm:write"), reamlController.CreateReaml)
	v1.GET("/reamls", middleware.ScopeMiddleware("realm:read"), reamlController.GetAllReamls)
	v1.GET("/reamls/:id", middleware.ScopeMiddleware("realm:read"), reamlController.GetReamlByID)

	userService := services.NewUserService(initilizer.DB)
	userController := controllers.NewUserController(userService)
	v1.GET("/users", userController.GetAllUsers)
	v1.POST("/users", userController.CreateUser)
	v1.GET("/users/:id", userController.GetUserByID)
	v1.POST("/signin", userController.SignIn)
	v1.POST("/authenticate", userController.AuthenticateUser)

	var reamlRoute = v1.Group(":reaml")
	reamlRoute.Use(utils.TokenMiddleware())

	// userService := services.NewUserService(initilizer.DB)

	// userController := controllers.NewUserController(userService)
	// v1.POST("/users", userController.CreateUser)

	// getUser := v1.Group("/users")
	// getUser.Use(utils.TokenMiddleware())
	// getUser.GET("/:id", userController.GetUserByID)

	// v1.POST("/users/:user_id/roles/:role_id", userController.AddRoleToUser)

	// v1.POST("/signin", userController.SignIn)

	// // utils.Resources(*v1, "users", controllers.NewUserController(initilizer.DB))
	// utils.Resources(*v1, "demos", controllers.NewDemoController(initilizer.DB))
	// // utils.Resources(*v1, "roles", &controllers.RoleController{DB: initilizer.DB})
	// // v1.POST("/users/:id/add_roles", controllers.NewUserController(initilizer.DB).AddRoles)

	// // r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
