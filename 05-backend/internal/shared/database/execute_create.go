package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// CreateOne performs a single insert operation
func CreateOne(ctx context.Context, db DatabaseConn, tableName string, columnNames []string, values []any) error {
	if ctx == nil || tableName == "" || len(columnNames) == 0 || len(values) == 0 {
		return fmt.Errorf("invalid arguments to execute insert")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	_, err = executor.CopyFrom(ctx, pgx.Identifier{tableName}, columnNames, pgx.CopyFromRows([][]any{values}))
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	return nil
}

// CreateMany performs a bulk insert operation
func CreateMany(ctx context.Context, db DatabaseConn, tableName string, columnNames []string, rows [][]any) error {
	if ctx == nil || tableName == "" || len(columnNames) == 0 || len(rows) == 0 {
		return fmt.Errorf("invalid arguments to execute bulk insert")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	_, err = executor.CopyFrom(ctx, pgx.Identifier{tableName}, columnNames, pgx.CopyFromRows(rows))

	if err != nil {
		return fmt.Errorf("bulk insert failed: %w", err)
	}

	return nil
}
