package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brajes/go-rest-api/internal/config"
	"github.com/brajes/go-rest-api/internal/initializer/database"
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
		fmt.Println("Failed to setup database")
	}

	fmt.Println(db.Config)

	e := echo.New()
	e.HideBanner = true
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
