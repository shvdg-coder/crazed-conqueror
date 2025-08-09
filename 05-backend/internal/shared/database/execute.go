package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Execute executes a script with optional arguments and returns no value
func Execute(ctx context.Context, executor Executor, script string, arguments ...any) error {
	if ctx == nil || script == "" {
		return fmt.Errorf("invalid arguments to execute script")
	}

	_, err := executor.Exec(ctx, script, arguments...)
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}

// Count executes a script with optional arguments and returns the number of rows affected
func Count(ctx context.Context, executor Executor, script string, arguments ...any) (int, error) {
	var count int
	if ctx == nil || script == "" {
		return count, fmt.Errorf("invalid arguments to count rows")
	}

	row := executor.QueryRow(ctx, script, arguments...)
	err := row.Scan(&count)
	if err != nil {
		return count, fmt.Errorf("failed to execute command: %w", err)
	}

	return count, nil
}

// QueryOne executes a script with arguments and returns a single scanned value
func QueryOne[T any](ctx context.Context, executor Executor, script string, arguments []any, scan ScannerFunc[T]) (T, error) {
	var zero T

	if ctx == nil || script == "" || scan == nil {
		return zero, fmt.Errorf("invalid arguments to query single result")
	}

	row := executor.QueryRow(ctx, script, arguments...)

	return scan(row)
}

// QueryMany executes a script with arguments and returns a slice of scanned values
func QueryMany[T any](ctx context.Context, executor Executor, script string, arguments []any, scan ScannerFunc[T]) ([]T, error) {
	var zero []T
	if ctx == nil || script == "" || scan == nil {
		return zero, fmt.Errorf("invalid arguments to query multiple results")
	}

	rows, err := executor.Query(ctx, script, arguments...)
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
}

// Batch executes multiple instances of the same script with different arguments
func Batch(ctx context.Context, executor Executor, script string, argumentSets [][]any) error {
	if ctx == nil || script == "" || len(argumentSets) == 0 {
		return fmt.Errorf("invalid arguments to execute batch")
	}

	batch := &pgx.Batch{}

	for _, arguments := range argumentSets {
		batch.Queue(script, arguments...)
	}

	results := executor.SendBatch(ctx, batch)
	err := results.Close()
	if err != nil {
		return fmt.Errorf("failed to execute batch: %w", err)
	}

	return nil
}
