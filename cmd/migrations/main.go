package main

import (
	"log"

	"github.com/bobchopperz/bahrululum/internal/config"
	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/init/database"
)

func main() {
	log.Println("Starting Migration")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Connecting", cfg.DatabaseConfig.Host)

	db, err := database.InitDatabase(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("Failed to setup database")
	}

	userModel := models.User{}

	db.AutoMigrate(&userModel)
}
