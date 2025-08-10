package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// UserCharacterSchema represents the user character schema operations.
type UserCharacterSchema struct {
	database.Connection
}

// NewUserCharacterSchema creates a new instance of UserCharacterSchema.
func NewUserCharacterSchema(connection database.Connection) *UserCharacterSchema {
	return &UserCharacterSchema{connection}
}

// CreateTable creates the user_characters table in the database
func (s *UserCharacterSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the user_characters table from the database
func (s *UserCharacterSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
