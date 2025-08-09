package contexts

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type txKey struct{}

// GetTransaction retrieves a transaction from context if it exists
func GetTransaction(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}

// SetTransaction adds a transaction to the context
func SetTransaction(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}
