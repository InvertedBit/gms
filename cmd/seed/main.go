package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/invertedbit/gms/auth"
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
	initUserEmail := os.Getenv("INIT_USER_EMAIL")
	initUserPw := os.Getenv("INIT_USER_PW")
	if dsn == "" {
		fmt.Println("DSN environment variable is required")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Role{}, &models.User{})

	database.DBConn = db

	// Create default roles
	adminRole := models.Role{
		Slug:        "admin",
		Name:        "Admin",
		Description: "Administrator with full access",
	}
	userRole := models.Role{
		Slug:        "user",
		Name:        "User",
		Description: "Regular user",
	}

	// Check if roles already exist, if not create them
	var existingAdmin models.Role
	result := db.Where("name = ?", "admin").First(&existingAdmin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if err := db.Create(&adminRole).Error; err != nil {
			fmt.Printf("Error creating admin role: %v\n", err)
		} else {
			fmt.Println("Created 'admin' role")
		}
	} else if result.Error != nil {
		fmt.Printf("Error checking for admin role: %v\n", result.Error)
	} else {
		fmt.Println("'admin' role already exists")
	}

	var existingUser models.Role
	result = db.Where("name = ?", "user").First(&existingUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if err := db.Create(&userRole).Error; err != nil {
			fmt.Printf("Error creating user role: %v\n", err)
		} else {
			fmt.Println("Created 'user' role")
		}
	} else if result.Error != nil {
		fmt.Printf("Error checking for user role: %v\n", result.Error)
	} else {
		fmt.Println("'user' role already exists")
	}

	hashedPassword, err := auth.HashPassword(initUserPw)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
		return
	}

	adminUser := models.User{
		Email:             initUserEmail,
		RoleSlug:          "admin",
		EncryptedPassword: hashedPassword,
	}

	// Check if admin user already exists
	var existingAdminUser models.User
	if err := db.Where("email = ?", adminUser.Email).First(&existingAdminUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&adminUser).Error; err != nil {
			fmt.Printf("Error creating admin user: %v\n", err)
		} else {
			fmt.Printf("Created admin user with email '%s'\n", adminUser.Email)
		}
	} else {
		fmt.Printf("Admin user with email '%s' already exists\n", existingAdminUser.Email)
	}

	fmt.Println("Seed completed successfully!")
}
