package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{Pool: pool}
}

func Initialize(conf *DatabaseConfig) (*pgxpool.Pool, error) {
	fmt.Println("Initializing database connection...")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	poolConfig, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		return nil, fmt.Errorf("error parsing pool config: %w", err)
	}

	poolConfig.MaxConns = int32(conf.PoolMaxConns)

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		return nil, fmt.Errorf("error establishing connection to the database: %w", err)
	}

	fmt.Println("Database connection established!")

	return pool, nil
}
