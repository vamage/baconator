// Purpose: This file is used to generate the api package from the openapi.yaml file.
package main

// This generates the api package from the openapi.yaml file
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target api -package api --clean openapi.yaml
