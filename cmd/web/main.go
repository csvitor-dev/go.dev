package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/src/apps/web/routes"
	"github.com/csvitor-dev/go.dev/src/router"
	"github.com/gorilla/mux"
)

func init() {
	env.LoadGeneralEnv()
	env.LoadWebEnv()
}

func main() {
	r := router.Generate(routes.All(),
		func(router *mux.Router) {
			fileServer := http.FileServer(http.Dir("./src/apps/web/static/"))
			router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
		})

	log.Printf("Listening on port '%s'\n", env.WebEnv.PORT)
	log.Fatalln(http.ListenAndServe(env.WebEnv.PORT, r))
}
