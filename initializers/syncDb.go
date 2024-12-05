package initializers

import (
	"authorization/models"
	"log"
)

func SyncDb() {
	var err error
	// Now you can use GORM features
	err = DB.AutoMigrate(&models.User{}) // Assuming 'User' is your model
	if err != nil {
		log.Fatal(err)
	}
}
