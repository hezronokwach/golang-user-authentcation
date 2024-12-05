package initializers

import (
	"log"

	"gorm.io/driver/sqlite" // Correct import for the SQLite driver
	"gorm.io/gorm"
)

var DB *gorm.DB // Declare DB as *gorm.DB

func ConnectToDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./app.db"), &gorm.Config{}) // Use gorm.Open
	if err != nil {
		log.Fatal(err)
	}
}
