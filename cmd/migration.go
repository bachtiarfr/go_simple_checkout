package main

import (
	"fmt"
	"log"
	"os"
	"simple-checkout-app/internal/config"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	config, err := config.ReadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Databasename,
	)

	migrationsDir := "internal/database/migrations"

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run migration.go <up|down>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "up":
		runMigrations(dbURL, migrationsDir, true)
	case "down":
		runMigrations(dbURL, migrationsDir, false)
	default:
		fmt.Println("Usage: go run migration.go <up|down>")
		os.Exit(1)
	}
}

func runMigrations(dbURL, migrationsDir string, up bool) {
	m, err := migrate.New(
		"file://"+migrationsDir,
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}

	if up {
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println("No migration needed.")
			} else {
				log.Fatalf("Failed to apply migrations: %v", err)
			}
		} else {
			fmt.Println("Migrations applied successfully.")
		}
	} else {
		if err := m.Steps(-1); err != nil {
			log.Fatalf("Failed to rollback migrations: %v", err)
		} else {
			fmt.Println("Migrations rolled back successfully.")
		}
	}
}
