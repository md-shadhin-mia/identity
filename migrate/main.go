package main

import (
	"indentity/initilizer"
	"indentity/models"
	"log"
)

func init() {
	initilizer.LoadEnv()
	initilizer.DBConnect()
}

func main() {

	if initilizer.DB == nil {
		log.Fatal("DB is nil")
	}

	initilizer.DB.AutoMigrate(&models.User{}, &models.Role{})
	// initilizer.DB.AutoMigrate()
	// initilizer.DB.AutoMigrate()
}
