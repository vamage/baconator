// Purpose: This file is used to generate any code for the module
package main

// This generates the api package from the openapi.yaml file
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target api -package api --clean openapi.yaml
// This generates the user module with sqlc
//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f sql/user/sqlc.yaml
