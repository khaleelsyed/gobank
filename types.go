package main

import (
	"math/rand/v2"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	AccountNumber int64     `json:"number"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:created_at`
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID:            rand.IntN(1000),
		FirstName:     firstName,
		LastName:      LastName,
		AccountNumber: rand.Int64(),
		CreatedAt:     time.Now().UTC(),
	}
}
