package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// SetupDatabaseConnection  is Creating a new connection database
func SetupDatabaseConnection () *gorm.DB{

	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser,dbPass,dbHost,dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create connection")
	}
	return db
}

//
func CloseConnectionDatabase(db *gorm.DB)  {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed close connection")
	}
	dbSQL.Close()
}
