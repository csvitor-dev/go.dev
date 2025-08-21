package config

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type WebEnviroment struct {
	PORT string
}

var WebEnv WebEnviroment

func LoadWebEnv() {
	var err error

	if err = env.Load(".env.web"); err != nil {
		log.Fatal(err)
	}
	WebEnv = WebEnviroment{
		PORT: os.Getenv("WEB_PORT"),
	}

	log.Println("Web enviroment variables loaded successfully!")
}
