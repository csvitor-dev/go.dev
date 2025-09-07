package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/csvitor-dev/go.dev/resources"
	"github.com/csvitor-dev/go.dev/src/apps/web/routes"
	"github.com/csvitor-dev/go.dev/src/router"
	"github.com/gorilla/mux"
)

func init() {
	env.LoadGeneralEnv()
	env.LoadWebEnv()

	err := resources.PrepareTailwind()

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := router.Generate(routes.All(),
		func(r *mux.Router) {
			fileServer := http.FileServer(http.Dir("./src/static/"))
			r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
		},
		func(r *mux.Router) {
			router.MapDefaultRoutes(r)
		},
	)

	log.Printf("Listening on port '%s'\n", env.WebEnv.PORT)
	log.Fatalln(http.ListenAndServe(env.WebEnv.PORT, r))
}
