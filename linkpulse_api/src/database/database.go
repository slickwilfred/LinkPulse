package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)

	if err != nil {
		return nil, fmt.Errorf("error establishing connection to the database: %w", err)
	}

	fmt.Println("Database connection established!")

	return pool, nil
}

// Checks to see if a user exists
func (db *DB) CheckUserExists(email string) (bool, error) {
	var exists bool
	err := db.Pool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	return exists, err
}

// Creates a new user
func (db *DB) CreateUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Pool.Exec(context.Background(), `INSERT INTO "users" (name, email, password) VALUES ($1, $2, $3)`, name, email, string(hashedPassword))
	return err
}
