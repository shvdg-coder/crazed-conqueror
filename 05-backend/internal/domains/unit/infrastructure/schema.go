package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// UnitSchema represents the unit schema operations.
type UnitSchema struct {
	database.Connection
}

// NewUnitSchema creates a new instance of UnitSchema.
func NewUnitSchema(connection database.Connection) *UnitSchema {
	return &UnitSchema{connection}
}

// CreateTable creates the units-table in the database
func (s *UnitSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the units-table from the database
func (s *UnitSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
