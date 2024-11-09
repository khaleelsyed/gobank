package main

import "math/rand/v2"

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber int64  `json:"number"`
	Balance       int64  `json:"balance"`
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID:            rand.IntN(1000),
		FirstName:     firstName,
		LastName:      LastName,
		AccountNumber: rand.Int64(),
	}
}
