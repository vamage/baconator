// Package main contains the main function that starts the server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vamage/baconator/handlers"
	"github.com/vamage/baconator/security"

	"github.com/vamage/baconator/api"
)

func main() {
	fmt.Println("Hello, World!")
	service := &handlers.Handler{}
	sec := &security.Security{}
	srv, _ := api.NewServer(service, sec)
	server := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           srv,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
