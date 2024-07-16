package main

import (
	"baconator/config"
	"log"
	"net/http"

	"baconator/api"
	"baconator/handlers"
	"baconator/security"
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
