package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/src/apps/web/routes"
	"github.com/csvitor-dev/social-media/src/router"
)

func init() {
	config.LoadWebEnv()
}

func main() {
	r := router.Generate(routes.All())

	log.Printf("Listening on port '%s'\n", config.WebEnv.PORT)
	log.Fatalln(http.ListenAndServe(config.WebEnv.PORT, r))
}
