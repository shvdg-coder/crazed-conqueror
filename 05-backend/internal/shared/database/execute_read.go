package database

import (
	"context"
	"fmt"
)

// ReadOne executes a script with arguments and returns a single value.
func ReadOne[T any](ctx context.Context, db DatabaseConn, script string, arguments []any, scan ScannerFunc[T]) (T, error) {
	var zero T

	if ctx == nil || script == "" || len(arguments) == 0 || scan == nil {
		return zero, fmt.Errorf("invalid arguments to execute read")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return zero, fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	row := executor.QueryRow(ctx, script, arguments...)

	return scan(row)
}

// ReadMany executes a script with arguments and returns a slice of values.
func ReadMany[T any](ctx context.Context, db DatabaseConn, script string, arguments []any, scan ScannerFunc[T]) ([]T, error) {
	var zero []T
	if ctx == nil || script == "" || len(arguments) == 0 || scan == nil {
		return zero, fmt.Errorf("invalid arguments to execute read")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return zero, fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

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
