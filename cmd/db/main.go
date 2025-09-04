package main

import (
	"log"
	"os"
	"strconv"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/internal/db/migrations"
)

func init() {
	env.LoadGeneralEnv()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide an argument: 'up', 'down', or 'status'")
		os.Exit(1)
	}
	log.Println("Running database migrations...")

	isForcedExec := len(os.Args) == 4 && os.Args[2] == "--force"
	var version int

	if isForcedExec {
		version = getVersionForPrompt()
	}

	if version == -1 {
		log.Println("Invalid version number:", version)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "up":
		if isForcedExec {
			forceMigration(version)
			os.Exit(0)
		}
		err := migrations.Up()

		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		log.Println("Migrations completed successfully.")
	case "down":
		if isForcedExec {
			forceMigration(version)
			os.Exit(0)
		}
		err := migrations.Down()

		if err != nil {
			log.Fatalln(err)
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

func getVersionForPrompt() int {
	version, err := strconv.Atoi(os.Args[3])

	if err != nil {
		return -1
	}
	return version
}

func forceMigration(version int) {
	err := migrations.Force(version)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Println("Force migration successfully")
}
