package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/internal/config"
	"github.com/csvitor-dev/social-media/src/apps/web/routes"
	"github.com/csvitor-dev/social-media/src/router"
	"github.com/gorilla/mux"
)

func init() {
	config.LoadGeneralEnv()
	config.LoadWebEnv()
}

func main() {
	r := router.Generate(routes.All(),
		func(router *mux.Router) {
			fileServer := http.FileServer(http.Dir("./src/apps/web/assets/"))
			router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
		})

	log.Printf("Listening on port '%s'\n", config.WebEnv.PORT)
	log.Fatalln(http.ListenAndServe(config.WebEnv.PORT, r))
}
