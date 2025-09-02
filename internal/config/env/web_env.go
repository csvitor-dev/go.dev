package env

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type WebEnvironment struct {
	PORT    string
	API_URL string
}

var WebEnv WebEnvironment

func LoadWebEnv() {
	var err error

	if err = env.Load(".env.web"); err != nil {
		log.Fatal(err)
	}
	WebEnv = WebEnvironment{
		PORT:    os.Getenv("WEB_PORT"),
		API_URL: os.Getenv("API_URL"),
	}

	log.Println("Web environment variables loaded successfully!")
}
