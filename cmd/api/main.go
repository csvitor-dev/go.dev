package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/src/apps/api/routes"
	"github.com/csvitor-dev/go.dev/src/router"
)

func init() {
	env.LoadGeneralEnv()
	env.LoadApiEnv()
	env.LoadEmailEnv()
}

func main() {
	r := router.Generate(routes.All())

	log.Printf("Listening on port '%s'\n", env.ApiEnv.PORT)
	log.Fatalln(http.ListenAndServe(env.ApiEnv.PORT, r))
}
