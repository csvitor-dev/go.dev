package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/src/api/router"
)

func init() {
	// Load environment variables
	config.LoadEnv()
}

func main() {
	r := router.Generate()

	log.Printf("Listening on port '%s'\n", config.Env.PORT)
	log.Fatalln(http.ListenAndServe(config.Env.PORT, r))
}
