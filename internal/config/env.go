package config

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type Enviroment struct {
	PORT              string
	CONNECTION_STRING string
}

var Env Enviroment

func LoadEnv() {
	var err error

	if err = env.Load(); err != nil {
		log.Fatal(err)
	}
	Env.PORT = os.Getenv("API_PORT")
	Env.CONNECTION_STRING = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
