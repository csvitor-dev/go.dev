package config

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

// SMTPConfig: SMTP server configuration
type SMTPEnviroment struct {
	EMAIL   string
	API_KEY string
}

var EmailEnv SMTPEnviroment

func LoadEmailEnv() {
	var err error

	if err = env.Load(".env.email"); err != nil {
		log.Fatal(err)
	}
	EmailEnv = SMTPEnviroment{
		EMAIL:   os.Getenv("EMAIL_SENDER"),
		API_KEY: os.Getenv("RESEND_API_KEY"),
	}

	log.Println("Email enviroment variables loaded successfully!")
}
