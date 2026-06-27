package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type DatabaseOptions struct {
	Driver          string
	Host            string
	User            string
	Password        string
	DBName          string
	Port            int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxConns        int32
	MinConns        int32
	Enabled         bool
	SSLMode         bool
}

func InitDB(log *zerolog.Logger, opt *DatabaseOptions) *pgxpool.Pool {
	if !opt.Enabled {
		return nil
	}

	config, err := pgxpool.ParseConfig(getURI((opt)))

	if err != nil {
		log.Panic().Err(err).Msg("failed to parse database config")
	}

	config.MaxConns = opt.MaxConns
	config.MinConns = opt.MinConns
	config.MaxConnLifetime = opt.ConnMaxLifetime
	config.MaxConnIdleTime = opt.ConnMaxIdleTime

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panic().
			Err(err).
			Str("driver", strings.ToUpper(opt.Driver)).
			Msg("database initialization failed")
	}

	log.Info().
		Str("driver", strings.ToUpper(opt.Driver)).
		Msg("database initialized")

	return pool
}

func getURI(opt *DatabaseOptions) string {
	ssl := "disable"

	if opt.SSLMode {
		ssl = "require"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		opt.User,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.DBName,
		ssl,
	)
}
