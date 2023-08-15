// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	// Create
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	// Delete
	DeleteAccount(ctx context.Context, accountID uuid.UUID) error
	// Read One
	GetAccount(ctx context.Context, accountID uuid.UUID) (Account, error)
	// Read Many
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	// Update
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
