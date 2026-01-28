package main

import (
	"os"

	"github.com/invertedbit/gms/database"
	"github.com/invertedbit/gms/handlers"
	"github.com/invertedbit/gms/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DSN string
}

func main() {

	godotenv.Load(".env.local")
	godotenv.Load()

	cfg := Config{
		DSN: os.Getenv("DSN"),
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate for auth schema models
	db.AutoMigrate(&models.Role{}, &models.User{}, &models.Media{}, &models.ComponentInstance{}, &models.ComponentProperty{}, &models.ComponentMedia{}, &models.Layout{}, &models.Page{})

	database.DBConn = db

	app := handlers.New()

	defer app.Shutdown()

	app.Listen(":3000")
}
