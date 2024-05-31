package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{Pool: pool}
}

func Initialize(conf *DatabaseConfig) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	fmt.Println("connecting with", connectionString)
	poolConfig, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		return nil, fmt.Errorf("error parsing pool config: %w", err)
	}

	poolConfig.MaxConns = int32(conf.PoolMaxConns)

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)

	if err != nil {
		return nil, fmt.Errorf("error establishing connection to the database: %w", err)
	}

	return pool, nil
}
