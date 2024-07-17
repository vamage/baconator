// Package main contains the main function that starts the server.
package main

import (
	"log"
	"net/http"

	"github.com/vamage/baconator/config"

	"github.com/vamage/baconator/api"
	"github.com/vamage/baconator/handlers"
	"github.com/vamage/baconator/security"
)

func main() {
	service := &handlers.Handler{}
	sec := &security.Security{}
	config.New()
	srv, err := api.NewServer(service, sec)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8081", srv); err != nil {
		log.Fatal(err)
	}
}
