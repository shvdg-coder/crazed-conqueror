package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Execute executes a script with optional arguments and returns no value
func Execute(ctx context.Context, connection Connection, script string, arguments ...any) error {
	if ctx == nil || script == "" {
		return fmt.Errorf("invalid arguments to execute script")
	}

	return WithExecutor(ctx, connection, func(executor Executor) error {
		_, err := executor.Exec(ctx, script, arguments...)
		if err != nil {
			return fmt.Errorf("failed to execute command: %w", err)
		}
		return nil
	})
}

// QueryOne executes a script with arguments and returns a single scanned value
func QueryOne[T any](ctx context.Context, connection Connection, query string, args []any, scan ScannerFunc[T]) (T, error) {
	var zero T

	if ctx == nil || query == "" || scan == nil {
		return zero, fmt.Errorf("invalid args to query single result")
	}

	return WithExecutorResult(ctx, connection, func(executor Executor) (T, error) {
		row := executor.QueryRow(ctx, query, args...)
		return scan(row)
	})
}

// QueryMany executes a script with arguments and returns a slice of scanned values
func QueryMany[T any](ctx context.Context, connection Connection, query string, args []any, scan ScannerFunc[T]) ([]T, error) {
	var zero []T
	if ctx == nil || query == "" || scan == nil {
		return zero, fmt.Errorf("invalid args to query multiple results")
	}

	return WithExecutorResult(ctx, connection, func(executor Executor) ([]T, error) {
		rows, err := executor.Query(ctx, query, args...)
		if err != nil {
			return zero, fmt.Errorf("failed to execute command: %w", err)
		}
		defer rows.Close()

		var results []T

		for rows.Next() {
			result, err := scan(rows)
			if err != nil {
				return zero, fmt.Errorf("failed to scan row: %w", err)
			}
			results = append(results, result)
		}

		return results, nil
	})
}

// Batch executes multiple instances of the same script with different arguments
func Batch(ctx context.Context, connection Connection, query string, argSets [][]any) error {
	if ctx == nil || query == "" || len(argSets) == 0 {
		return fmt.Errorf("invalid arguments to execute batch")
	}

	return WithExecutor(ctx, connection, func(executor Executor) error {
		batch := &pgx.Batch{}

		for _, args := range argSets {
			batch.Queue(query, args...)
		}

		results := executor.SendBatch(ctx, batch)
		err := results.Close()
		if err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}

		return nil
	})
}
