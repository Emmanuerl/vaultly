package config

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// connectDB establishes a connection to a postgres Server and wraps this connection
// around a bunDB instance, which is then returned
func ConnectDB(env *AppEnv) (*bun.DB, error) {
	config, err := pgx.ParseConfig(env.PostgresUrl)
	if err != nil {
		return nil, err
	}
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	return db, nil
}
