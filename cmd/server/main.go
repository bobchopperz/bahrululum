package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bobchopperz/bahrululum/internal/api"
	"github.com/bobchopperz/bahrululum/internal/config"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/init/database"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting server")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDatabase(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("Failed to setup database")
	}

	e := echo.New()
	e.HideBanner = true

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	api.SetupRoutes(e, userService)

	go func() {
		addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
		fmt.Println("Starting server", addr)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			fmt.Println("error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown")
	}

	fmt.Println("Server exited")
}
