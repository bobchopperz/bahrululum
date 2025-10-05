package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	mymiddleware "github.com/bobchopperz/bahrululum/internal/api/middleware"
	"github.com/bobchopperz/bahrululum/internal/api/routes"
	"github.com/bobchopperz/bahrululum/internal/api/validators"
	"github.com/bobchopperz/bahrululum/internal/config"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
	"github.com/bobchopperz/bahrululum/internal/domain/service"
	"github.com/bobchopperz/bahrululum/internal/init/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDatabase(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("Failed to setup database")
	}

	e := echo.New()
	e.Validator = validators.NewValidator()
	e.HideBanner = true
	configureMiddleware(e)

	userRepository := repository.NewUserRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	enrollmentRepository := repository.NewEnrollmentRepository(db)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository, &cfg.JWTConfig)
	courseService := service.NewCourseService(courseRepository)
	enrollmentService := service.NewEnrollmentService(enrollmentRepository)

	routes.SetupHealthRoutes(e)

	opts := routes.AuthRoutesOpts{
		AuthService: authService,
		UserService: userService,
	}
	routes.SetupAuthRoutes(e, opts)
	routes.SetupUsersRoutes(e, userService)
	routes.SetupCoursesRoutes(e, courseService)
	routes.SetupEnrollmentRoutes(e, enrollmentService)

	startServer(e, cfg)
}

func configureMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mymiddleware.CORS())
	e.Use(middleware.RequestID())
}

func startServer(e *echo.Echo, cfg *config.Config) {
	log.Println("Starting server")

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
