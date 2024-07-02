package main

import (
	"log"
	"net/http"

	"baconator/api"
	"baconator/handlers"
	"baconator/security"
)

func main() {
	service := &handlers.Handler{}
	sec := &security.Security{}
	srv, err := api.NewServer(service, sec)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
