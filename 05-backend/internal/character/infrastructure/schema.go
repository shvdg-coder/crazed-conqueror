package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// CharacterSchema represents the character schema operations.
type CharacterSchema struct {
	database.Connection
}

// NewCharacterSchema creates a new instance of CharacterSchema.
func NewCharacterSchema(connection database.Connection) *CharacterSchema {
	return &CharacterSchema{connection}
}

// CreateTable creates the characters-table in the database
func (s *CharacterSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the characters-table from the database
func (s *CharacterSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
