package config

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type ApiEnvironment struct {
	PORT    string
	WEB_URL string
}

var ApiEnv ApiEnvironment

func LoadApiEnv() {
	var err error

	if err = env.Load(".env.api"); err != nil {
		log.Fatal(err)
	}
	ApiEnv = ApiEnvironment{
		PORT:    os.Getenv("API_PORT"),
		WEB_URL: os.Getenv("WEB_URL"),
	}

	log.Println("API environment variables loaded successfully!")
}
