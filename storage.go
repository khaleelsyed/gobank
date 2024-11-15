package main

import "database/sql"

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccount(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgresStore, error) {
	return nil, nil
}
