package database

import (
	"context"
	"fmt"
)

// ExecuteRaw executes a script with optional arguments and returns no value
func ExecuteRaw(ctx context.Context, db DatabaseConn, script string, arguments ...any) error {
	if ctx == nil || script == "" {
		return fmt.Errorf("invalid arguments to execute script")
	}

	executor, cleanup, err := db.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	_, err = executor.Exec(ctx, script, arguments...)
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
