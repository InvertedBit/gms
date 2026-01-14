package main

import (
	"fmt"
	"os"

	"github.com/invertedbit/gms/database"
	"github.com/invertedbit/gms/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	dsn := os.Getenv("DSN")
	if dsn == "" {
		fmt.Println("DSN environment variable is required")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.DBConn = db

	// Create default roles
	adminRole := models.Role{
		Name:        "admin",
		Description: "Administrator with full access",
	}
	userRole := models.Role{
		Name:        "user",
		Description: "Regular user",
	}

	// Check if roles already exist, if not create them
	var existingAdmin models.Role
	result := db.Where("name = ?", "admin").First(&existingAdmin)
	if result.Error != nil {
		if err := db.Create(&adminRole).Error; err != nil {
			fmt.Printf("Error creating admin role: %v\n", err)
		} else {
			fmt.Println("Created 'admin' role")
		}
	} else {
		fmt.Println("'admin' role already exists")
	}

	var existingUser models.Role
	result = db.Where("name = ?", "user").First(&existingUser)
	if result.Error != nil {
		if err := db.Create(&userRole).Error; err != nil {
			fmt.Printf("Error creating user role: %v\n", err)
		} else {
			fmt.Println("Created 'user' role")
		}
	} else {
		fmt.Println("'user' role already exists")
	}

	fmt.Println("Seed completed successfully!")
}
