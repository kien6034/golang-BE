package db

import (
	"context"
	"database/sql"
	"fmt"
)

// // Store defines all functions to execute db queries and transactions
// type Store interface {
// 	Querier
// 	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
// }

// SQLStore provides all functions to execute SQL queries and transactions
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {

	return &Store{
		db:      db,
		Queries: New(db),
	}
}

var txKey = struct{}{}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
