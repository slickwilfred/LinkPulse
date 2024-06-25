package models

import (
	"context"
	"database/sql"
	"fmt"
	"tracking_backend/src/database"
	"tracking_backend/src/dtos"

	"github.com/georgysavva/scany/v2/pgxscan"
	"golang.org/x/crypto/bcrypt"
)

type User_Model struct {
	DB   *database.DB
	User *dtos.User
}

func NewUserModel(db *database.DB) *User_Model {
	return &User_Model{DB: db, User: &dtos.User{}}
}
func (um *User_Model) LoginUser(req dtos.LoginRequest) (*dtos.User, error) {
	user, err := um.GetUserByEmail(req.Email)

	if err != nil {
		// Check if the error indicated that no rows were found
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is no account associated with the email: " + req.Email)
		}

		// For any other errors, return it
		return nil, err
	}

	// Compare password and hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Passwords don't match
		return nil, fmt.Errorf("invalid password")
	}

	// User is authenticated
	return user, nil
}

func (um *User_Model) GetUserByEmail(email string) (*dtos.User, error) {
	var user dtos.User
	query := "SELECT FROM Users WHERE email = $1"
	err := pgxscan.Get(context.Background(), um.DB.Pool, &user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Checks to see if a user exists
func (um *User_Model) CheckUserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := pgxscan.Get(context.Background(), um.DB.Pool, exists, query, email)
	return exists, err
}

// Creates a new user
func (um *User_Model) CreateUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := "INSERT INTO users VALUES (name, email, password)"
	_, err = um.DB.Pool.Exec(context.Background(), query, name, email, string(hashedPassword))
	return err
}
