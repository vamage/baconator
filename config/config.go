package config

import (
	"context"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
)

type Configuration struct {
	Database string
	Pool     *pgxpool.Pool
	Logger   *slog.Logger
}

var conf *Configuration

func New() *Configuration {
	pool, err := pgxpool.ParseConfig("postgresql://postgres:example@db/baconator")
	if err != nil {
		panic(err)
	}
	pool.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	npool, err := pgxpool.NewWithConfig(context.Background(), pool)
	conf = &Configuration{
		Database: "baconator",
		Pool:     npool,
		Logger:   l,
	}
	return conf
}
func GetConf() *Configuration {
	return conf

}
