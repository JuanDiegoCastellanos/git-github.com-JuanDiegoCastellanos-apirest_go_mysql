// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Account struct {
	ID        int64     `db:"id" json:"id"`
	Owner     string    `db:"owner" json:"owner"`
	Balance   int64     `db:"balance" json:"balance"`
	Currency  string    `db:"currency" json:"currency"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Entry struct {
	ID        int64 `db:"id" json:"id"`
	AccountID int64 `db:"account_id" json:"account_id"`
	// can be negative or positive
	Amount    int64     `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Transfer struct {
	ID            int64 `db:"id" json:"id"`
	FromAccountID int64 `db:"from_account_id" json:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id" json:"to_account_id"`
	// must be positive
	Amount    int64     `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
