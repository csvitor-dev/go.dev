package config

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

// SMTPConfig: SMTP server configuration
type SMTPEnviroment struct {
	HOST     string
	PORT     string
	EMAIL    string
	USERNAME string
	API_KEY  string
}

var EmailEnv SMTPEnviroment

func LoadEmailEnv() {
	var err error

	if err = env.Load(".env.email"); err != nil {
		log.Fatal(err)
	}
	EmailEnv = SMTPEnviroment{
		HOST:     os.Getenv("EMAIL_SMTP"),
		PORT:     os.Getenv("EMAIL_PORT"),
		EMAIL:    os.Getenv("EMAIL_SENDER"),
		USERNAME: os.Getenv("EMAIL_USER"),
		API_KEY:  os.Getenv("RESEND_API_KEY"),
	}
	EmailEnv.EMAIL = fmt.Sprintf("%s <%s>", EmailEnv.USERNAME, EmailEnv.EMAIL)

	log.Println("Email enviroment variables loaded successfully")
}
