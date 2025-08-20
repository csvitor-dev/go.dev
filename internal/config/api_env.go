package config

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type ApiEnviroment struct {
	SECRET_KEY        []byte
	PORT              string
	CONNECTION_STRING string
}

var ApiEnv ApiEnviroment

func LoadApiEnv() {
	var err error

	if err = env.Load(".env"); err != nil {
		log.Fatal(err)
	}
	ApiEnv = ApiEnviroment{
		SECRET_KEY: []byte(os.Getenv("AUTH_SECRET_KEY")),
		PORT:       os.Getenv("API_PORT"),
		CONNECTION_STRING: fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
	}
	log.Println("API enviroment variables loaded successfully")
}
