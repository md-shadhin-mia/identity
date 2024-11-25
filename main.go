package main

import (
	"flag"
	"indentity/controllers"
	"indentity/core"
	"indentity/core/cache"
	"indentity/initilizer"
	"indentity/middleware"
	"indentity/migrate"
	"indentity/services"
	"indentity/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizer.LoadEnv()
	initilizer.DBConnect()
	cache.InitializeCache()
}
func serve() {
	r := core.NewEngine()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("assets/templates/*.html")
	var v1 = r.Group("/" + os.Getenv("Name"))

	reamlService := services.NewRealmService(initilizer.DB)
	reamlController := controllers.NewReamlController(reamlService)
	v1.POST("/reamls", middleware.ScopeMiddleware(utils.RealmWrite), reamlController.CreateReaml)
	v1.GET("/reamls", middleware.ScopeMiddleware(utils.RealmRead), reamlController.GetAllReamls)
	v1.GET("/reamls/:id", middleware.ScopeMiddleware(utils.RealmRead), reamlController.GetReamlByID)
	v1.PUT("/reamls/:id", middleware.ScopeMiddleware(utils.RealmUpdate), reamlController.UpdateReaml)
	v1.DELETE("/reamls/:id", middleware.ScopeMiddleware(utils.RealmDelete), reamlController.DeleteReaml)

	clientService := services.NewClientService(initilizer.DB)
	clientController := controllers.NewClientController(clientService)
	v1.GET("/clients-by-name/:name", middleware.ScopeMiddleware(utils.ClientRead), clientController.GetClientByName)
	v1.POST("/clients", middleware.ScopeMiddleware(utils.ClientWrite), clientController.CreateClient)
	v1.GET("/clients/:id", middleware.ScopeMiddleware(utils.ClientRead), clientController.GetClientByID)
	v1.PUT("/clients/:id", middleware.ScopeMiddleware(utils.ClientUpdate), clientController.UpdateClient)
	v1.DELETE("/clients/:id", middleware.ScopeMiddleware(utils.ClientDelete), clientController.DeleteClient)

	userService := services.NewUserService(initilizer.DB)
	userController := controllers.NewUserController(userService)
	v1.GET("/users", userController.GetAllUsers)
	v1.POST("/users", userController.CreateUser)
	v1.GET("/users/:id", userController.GetUserByID)
	v1.POST("/signin", userController.SignIn)
	v1.POST("/authenticate", userController.AuthenticateUser)
	authorizeService := services.NewAuthorizeService(initilizer.DB)
	authorizeController := controllers.NewAuthorizeController(authorizeService)
	v1.GET("/authorize", authorizeController.Authorize)
	v1.GET("/token", authorizeController.TokenAuthorize)

	var reamlRoute = v1.Group(":reaml")
	reamlRoute.Use(utils.TokenMiddleware())

	r.Run() // listen and serve on 0.0.0.0:8080
}
func main() {
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	serveCmd.Usage = func() {
		println("Usage: ./app serve [options]")
	}

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		serveCmd.Parse(os.Args[2:])
		serve()
	} else if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrate.Migrate()
	} else {
		println("Usage: ./app serve [options] or ./app migrate")
	}
}
