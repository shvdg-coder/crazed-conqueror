package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// FormationSchema represents the formation schema operations.
type FormationSchema struct {
	database.Connection
}

// NewFormationSchema creates a new instance of FormationSchema.
func NewFormationSchema(connection database.Connection) *FormationSchema {
	return &FormationSchema{connection}
}

// CreateTable creates the formations table in the database with JSONB support
func (s *FormationSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the formations table from the database
func (s *FormationSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
