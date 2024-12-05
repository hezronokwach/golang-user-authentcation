package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dbPath := "./db2.db"

	// Delete the existing database file
	// if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
	//     log.Printf("Error removing existing database: %v", err)
	// }

	// Create fresh database connection
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to fresh database")
}
