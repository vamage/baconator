// Package config contains the Configuration for application and the shared objects.
package config

import (
	"context"
	"log/slog"
	"os"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Configuration is a struct that contains the shared resources for the application.
type Configuration struct {
	Database string
	Pool     *pgxpool.Pool
	Logger   *slog.Logger
}

// conf is a global variable that holds the Configuration struct.
var conf *Configuration

// New creates a new Configuration struct.
func New() *Configuration {
	pgurl := os.Getenv("DATABASE_URL")
	if pgurl == "" {
		pgurl = "postgresql://postgres:example@127.0.0.1/baconator"
	}
	pool, err := pgxpool.ParseConfig(pgurl)
	if err != nil {
		panic(err)
	}
	pool.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	npool, err := pgxpool.NewWithConfig(context.Background(), pool)
	if err != nil {
		panic(err)
	}
	conf = &Configuration{
		Database: "baconator",
		Pool:     npool,
		Logger:   l,
	}
	return conf
}

// GetConf returns the Configuration struct.
func GetConf() *Configuration {
	return conf
}
