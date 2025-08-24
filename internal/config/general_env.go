package config

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type Environment struct {
	SECRET_KEY        []byte
	CONNECTION_STRING string
}

var Env Environment

func LoadGeneralEnv() {
	var err error

	if err = env.Load(".env"); err != nil {
		log.Fatal(err)
	}
	Env = Environment{
		SECRET_KEY: []byte(os.Getenv("AUTH_SECRET_KEY")),
		CONNECTION_STRING: fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
	}

	log.Println("General environment variables loaded successfully!")
}
