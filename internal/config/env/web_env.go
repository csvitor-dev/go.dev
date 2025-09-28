package env

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type WebEnvironment struct {
	PORT      string
	API_URL   string
	HASH_KEY  []byte
	BLOCK_KEY []byte
}

var WebEnv WebEnvironment

func LoadWebEnv() {
	var err error

	if err = env.Load(".env.web"); err != nil {
		log.Fatal(err)
	}
	WebEnv = WebEnvironment{
		PORT:      os.Getenv("WEB_PORT"),
		API_URL:   os.Getenv("API_URL"),
		HASH_KEY:  []byte(os.Getenv("HASH_KEY")),
		BLOCK_KEY: []byte(os.Getenv("BLOCK_KEY")),
	}

	log.Println("Web environment variables loaded successfully!")
}
