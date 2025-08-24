package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/src/apps/api/routes"
	"github.com/csvitor-dev/social-media/src/router"
)

func init() {
	config.LoadGeneralEnv()
	config.LoadApiEnv()
	config.LoadEmailEnv()
}

func main() {
	r := router.Generate(routes.All())

	log.Printf("Listening on port '%s'\n", config.ApiEnv.PORT)
	log.Fatalln(http.ListenAndServe(config.ApiEnv.PORT, r))
}
