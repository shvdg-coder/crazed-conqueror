package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// CharacterUnitSchema represents the character unit schema operations.
type CharacterUnitSchema struct {
	database.Connection
}

// NewCharacterUnitSchema creates a new instance of CharacterUnitSchema.
func NewCharacterUnitSchema(connection database.Connection) *CharacterUnitSchema {
	return &CharacterUnitSchema{connection}
}

// CreateTable creates the character_units table in the database
func (s *CharacterUnitSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the character_units table from the database
func (s *CharacterUnitSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
