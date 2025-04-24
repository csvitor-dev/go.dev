package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/src/api/router"
)

func main() {
	config.LoadEnv()

	r := router.Generate()
	log.Printf("Listening on port '%s'\n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r))
}
