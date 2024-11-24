package migrate

import (
	"indentity/initilizer"
	"indentity/models"
	"log"

	"github.com/google/uuid"
)

func Migrate() {

	if initilizer.DB == nil {
		log.Fatal("DB is nil")
	}

	// initilizer.DB.AutoMigrate(&models.User{}, &models.Role{})
	// initilizer.DB.AutoMigrate()
	// initilizer.DB.AutoMigrate()
	initilizer.DB.AutoMigrate(&models.Realm{}, &models.User{}, &models.AdministrativeUser{})
	id := uuid.MustParse("1ddea550-c59d-4e0a-a2fc-4482982c383c")
	//master realm create
	masterRealm := models.Realm{
		Name:        "master",
		Description: "The master realm",
	}
	masterRealm.ID = id

	initilizer.DB.Save(&masterRealm)

	//admin user create
	adminUser := models.User{
		Username: "admin",
		Email:    "admin@localhost",
		RealmID:  masterRealm.ID,
	}
	adminUser.SetPassword("12345678")
	initilizer.DB.FirstOrCreate(&adminUser, models.User{Username: "admin"})

	/*	userList := make([]models.User, 0, 100)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i += 100 {
				userList = userList[:0]
				for j := i; j < i+100; j++ {
					username := fmt.Sprintf("user-%d", j)
					email := fmt.Sprintf("%s@localhost", username)
					password := fmt.Sprintf("$2a$12$%s", strings.Repeat("a", 12))
					userList = append(userList, models.User{
						Username: username,
						Email:    email,
						Password: password,
						RealmID:  masterRealm.ID,
					})
				}
				initilizer.DB.Create(&userList)
			}
		}()
		wg.Wait() */

	user := models.User{}
	initilizer.DB.First(&user, models.User{
		Username: "admin",
	})
	log.Printf("user: %+v", user)
	administrative := models.AdministrativeUser{
		UserId: user.ID,
	}
	initilizer.DB.FirstOrCreate(&administrative, models.AdministrativeUser{UserId: user.ID})
}
