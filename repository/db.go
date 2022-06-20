package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func New(user, pwd, nameDB string) (*DB, error) {
	db, err := connect(user, pwd, nameDB)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func connect(user, pwd, nameDB string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=5432 sslmode=disable", user, pwd, nameDB)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
