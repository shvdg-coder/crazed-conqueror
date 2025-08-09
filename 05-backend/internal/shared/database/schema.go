package database

import (
	"context"
	"fmt"
)

// DomainSchema defines the interface for database schema creation and deletion.
type DomainSchema interface {
	CreateTables(ctx context.Context, executor Executor) error
	DropTables(ctx context.Context, executor Executor) error
}

// Schemas manages database schemas.
type Schemas struct {
	database Connection
	schemas  []DomainSchema
}

// NewSchemas creates a new instance of Schemas.
func NewSchemas(database Connection, schemas ...DomainSchema) *Schemas {
	return &Schemas{
		database: database,
		schemas:  schemas,
	}
}

// CreateAllTables creates all tables in the database.
func (sm *Schemas) CreateAllTables(ctx context.Context) error {
	executor, cleanup, err := sm.database.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	for _, schema := range sm.schemas {
		if err = schema.CreateTables(ctx, executor); err != nil {
			return fmt.Errorf("failed to create tables: %w", err)
		}
	}

	return nil
}

// DropAllTables drops all tables in the database, in reverse order of creation.
func (sm *Schemas) DropAllTables(ctx context.Context) error {
	executor, cleanup, err := sm.database.GetExecutor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get executor: %w", err)
	}
	defer cleanup()

	for i := len(sm.schemas) - 1; i >= 0; i-- {
		if err = sm.schemas[i].DropTables(ctx, executor); err != nil {
			return fmt.Errorf("failed to drop tables: %w", err)
		}
	}

	return nil
}
