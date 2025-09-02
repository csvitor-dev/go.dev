package env

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

// EmailEnvironment: SMTP server configuration
type EmailEnvironment struct {
	EMAIL   string
	API_KEY string
}

var EmailEnv EmailEnvironment

func LoadEmailEnv() {
	var err error

	if err = env.Load(".env.email"); err != nil {
		log.Fatal(err)
	}
	EmailEnv = EmailEnvironment{
		EMAIL:   os.Getenv("EMAIL_SENDER"),
		API_KEY: os.Getenv("RESEND_API_KEY"),
	}

	log.Println("Email environment variables loaded successfully!")
}
