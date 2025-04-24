package config

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

var (
	ConnectionString string = ""
	Port 		     string = ""
)

func LoadEnv() {
	var err error

	if err = env.Load(); err != nil {
		log.Fatal(err)
	}
	Port = os.Getenv("API_PORT")
	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
