package database

import (
	"context"
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// Delete performs a single delete operation
func Delete(ctx context.Context, db Connection, tableName string, where string, arguments []any) error {
	if ctx == nil || tableName == "" || where == "" {
		return fmt.Errorf("invalid arguments to execute delete")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	_, err = executor.Exec(ctx, sql.CreateDeleteQuery(tableName, where), arguments...)
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
