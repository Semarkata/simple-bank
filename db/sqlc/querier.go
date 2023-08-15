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
	// Create
	CreateEntrie(ctx context.Context, arg CreateEntrieParams) (Entry, error)
	// Create
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	// Delete
	DeleteAccount(ctx context.Context, accountID uuid.UUID) error
	// Read One
	GetAccount(ctx context.Context, accountID uuid.UUID) (Account, error)
	// Read One
	GetEntry(ctx context.Context, entrieID uuid.UUID) (Entry, error)
	// Read One
	GetTransfer(ctx context.Context, transferID uuid.UUID) (Transfer, error)
	// Read Many
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	// Read Many
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	// Read Many
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	// Update
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
