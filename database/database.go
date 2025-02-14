package database

import (
	"database/sql"
	"log"
	"todo-app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // Import the modernc.org/sqlite package
)

var DB *gorm.DB

func InitDB() {
	// Open a connection to the SQLite database using modernc.org/sqlite
	sqlDB, err := sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Create a GORM DB instance using the sql.DB connection
	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to create GORM DB instance:", err)
	}

	// Auto migrate models
	err = DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("❌ Failed to auto migrate models:", err)
	}

	log.Println("✅ Database connected successfully")
}