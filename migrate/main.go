package main

import (
	"indentity/initilizer"
	"indentity/models"
	"log"

	"github.com/google/uuid"
)

func init() {
	initilizer.LoadEnv()
	initilizer.DBConnect()
}

func main() {

	if initilizer.DB == nil {
		log.Fatal("DB is nil")
	}

	// initilizer.DB.AutoMigrate(&models.User{}, &models.Role{})
	// initilizer.DB.AutoMigrate()
	// initilizer.DB.AutoMigrate()
	initilizer.DB.AutoMigrate(&models.Realm{}, &models.User{})
	id := uuid.MustParse("1ddea550-c59d-4e0a-a2fc-4482982c383c")
	//master realm create
	masterRealm := models.Realm{
		Name:        "master",
		Description: "The master realm",
	}
	masterRealm.ID = id

	initilizer.DB.Find(&masterRealm, "id = ?", id).FirstOrCreate(&masterRealm)

	//admin user create
	adminUser := models.User{
		Username: "admin",
		Email:    "admin@localhost",
		Password: "$2a$12$ZhxfLJXtjhMAuNVND.VOkeovlaGEJwAODoR7u.xXJdz/3ZewRYvLS",
		RealmID:  masterRealm.ID,
	}
	initilizer.DB.Create(&adminUser)
}
