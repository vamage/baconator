//go:build tools

package main

// Import to keep it in go.mod.
import (
	_ "github.com/ogen-go/ogen/cmd/ogen"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
