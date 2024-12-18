package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgresStore, error) {
	connStr := fmt.Sprintf("user=postgres dbname=postgres password=%s port=%s sslmode=disable", os.Getenv("POSTGRES_PASSWORD"), os.Getenv("APP_POSTGRES_PORT"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account(
	id SERIAL PRIMARY KEY,
	first_name varchar(50),
	last_name varchar(50),
	number text,
	balance serial,
	created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `
	INSERT INTO account
	(first_name, last_name, number, balance, created_at)
	VALUES (
	$1, $2, $3, $4, $5
	)`

	resp, err := s.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.AccountNumber,
		account.Balance,
		account.CreatedAt,
	)

	if err != nil {
		return err
	}

	log.Println("&+v\n", resp)

	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	response, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for response.Next() {
		return scanAccount(response)
	}
	return nil, fmt.Errorf("Account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scanAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}
