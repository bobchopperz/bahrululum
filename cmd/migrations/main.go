package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bobchopperz/bahrululum/internal/config"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

var (
	flags   = flag.NewFlagSet("goose", flag.ExitOnError)
	dir     = flags.String("dir", "./migrations", "directory with migration files")
	verbose = flags.Bool("v", false, "enable verbose mode")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		printUsage()
		return
	}

	command := args[0]

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if command == "create-db" {
		createDatabase(cfg)
		return
	}
	if command == "drop-db" {
		dropDatabase(cfg)
		return
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.SSLMode,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	if *verbose {
		goose.SetVerbose(true)
	}

	arguments := []string{}
	if len(args) > 1 {
		arguments = args[1:]
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

func printUsage() {
	fmt.Println("Usage: go run cmd/migrations/main.go [command] [args...]")
	fmt.Println()
	fmt.Println("Database Commands:")
	fmt.Println("  create-db          Create the database")
	fmt.Println("  drop-db            Drop the database")
	fmt.Println()
	fmt.Println("Migration Commands:")
	fmt.Println("  up                 Migrate the DB to the most recent version available")
	fmt.Println("  down               Roll back the version by 1")
	fmt.Println("  reset              Roll back all migrations")
	fmt.Println("  status             Print the status of all migrations")
	fmt.Println("  version            Print the current version of the database")
	fmt.Println("  create NAME [sql]  Creates new migration file with the current timestamp")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -dir string        directory with migration files (default \"./migrations\")")
	fmt.Println("  -v                 enable verbose mode")
}

func createDatabase(cfg *config.Config) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=%s",
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.SSLMode,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}
	defer db.Close()

	var exists bool
	query := "SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)"
	err = db.QueryRow(query, cfg.DatabaseConfig.Name).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	if exists {
		log.Printf("Database '%s' already exists", cfg.DatabaseConfig.Name)
		return
	}

	createQuery := fmt.Sprintf("CREATE DATABASE %s", cfg.DatabaseConfig.Name)
	_, err = db.Exec(createQuery)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	log.Printf("Database '%s' created successfully", cfg.DatabaseConfig.Name)
}

func dropDatabase(cfg *config.Config) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=%s",
		cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.SSLMode,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}
	defer db.Close()

	var exists bool
	query := "SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)"
	err = db.QueryRow(query, cfg.DatabaseConfig.Name).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	if !exists {
		log.Printf("Database '%s' does not exist", cfg.DatabaseConfig.Name)
		return
	}

	terminateQuery := fmt.Sprintf(`
		SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = '%s'
		  AND pid <> pg_backend_pid()
	`, cfg.DatabaseConfig.Name)

	_, err = db.Exec(terminateQuery)
	if err != nil {
		log.Printf("Warning: Failed to terminate connections: %v", err)
	}

	dropQuery := fmt.Sprintf("DROP DATABASE %s", cfg.DatabaseConfig.Name)
	_, err = db.Exec(dropQuery)
	if err != nil {
		log.Fatalf("Failed to drop database: %v", err)
	}

	log.Printf("Database '%s' dropped successfully", cfg.DatabaseConfig.Name)
}
