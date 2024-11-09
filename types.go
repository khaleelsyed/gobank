package main

import "math/rand/v2"

type Account struct {
	ID            int
	FirstName     string
	LastName      string
	AccountNumber int64
	Balance       int64
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID:            rand.IntN(1000),
		FirstName:     firstName,
		LastName:      LastName,
		AccountNumber: rand.Int64(),
	}
}
