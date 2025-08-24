package main

import (
	"log"
	"os"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/internal/db/migrations"
)

func init() {
	config.LoadGeneralEnv()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide an argument: 'up', 'down', or 'status'")
		os.Exit(1)
	}
	log.Println("Running database migrations...")

	switch os.Args[1] {
	case "up":
		err := migrations.Up()

		if err != nil {
			log.Fatalln("Failed to apply migrations:", err)
			os.Exit(1)
		}
		log.Println("Migrations completed successfully.")
	case "down":
		err := migrations.Down()

		if err != nil {
			log.Fatalln("Failed to revert migrations:", err)
			os.Exit(1)
		}
		log.Println("Migrations reverted successfully.")
	case "status":
		status, err := migrations.Status()

		if err != nil {
			log.Fatalln("Failed to get migration version:", err)
		}
		log.Println("Migration status:", status)
	default:
		log.Fatalln("Invalid argument. Use 'up' to run migrations or 'down' to revert them.")
		os.Exit(1)
	}
}
