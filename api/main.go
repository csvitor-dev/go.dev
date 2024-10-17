package main

import (
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/api/src/router"
)

func main() {
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}