package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/api/src/config"
	"github.com/csvitor-dev/social-media/api/src/router"
)

func main() {
	config.LoadEnv()

	r := router.Generate()
	log.Printf("Listening on port '%s'\n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r))
}
