package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// UpdateOne executes an update query and returns the updated record
func UpdateOne[T any](ctx context.Context, db DatabaseConn, script string, arguments []any, scan ScannerFunc[T]) (T, error) {
	var zero T

	if ctx == nil || script == "" || len(arguments) == 0 || scan == nil {
		return zero, fmt.Errorf("invalid arguments to execute update")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return zero, fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	row := executor.QueryRow(ctx, script, arguments...)

	return scan(row)
}

// UpdateMany executes multiple update queries using a batch
func UpdateMany(ctx context.Context, db DatabaseConn, script string, rows [][]any) error {
	if ctx == nil || script == "" || len(rows) == 0 {
		return fmt.Errorf("invalid arguments to execute batch update")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	batch := &pgx.Batch{}

	for _, arguments := range rows {
		batch.Queue(script, arguments...)
	}

	results := executor.SendBatch(ctx, batch)
	err = results.Close()
	if err != nil {
		return fmt.Errorf("failed to execute batch update: %w", err)
	}

	return nil
}
