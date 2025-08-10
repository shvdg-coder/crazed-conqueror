package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// CharacterFormationSchema represents the character formation schema operations.
type CharacterFormationSchema struct {
	database.Connection
}

// NewCharacterFormationSchema creates a new instance of CharacterFormationSchema.
func NewCharacterFormationSchema(connection database.Connection) *CharacterFormationSchema {
	return &CharacterFormationSchema{connection}
}

// CreateTable creates the character_formations table in the database
func (s *CharacterFormationSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the character_formations table from the database
func (s *CharacterFormationSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
