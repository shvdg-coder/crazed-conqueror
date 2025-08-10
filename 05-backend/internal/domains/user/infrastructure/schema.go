package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// UserSchema represents the user schema operations.
type UserSchema struct {
	database.Connection
}

// NewUserSchema creates a new instance of UserSchema.
func NewUserSchema(connection database.Connection) *UserSchema {
	return &UserSchema{connection}
}

// CreateTable creates the users-table in the database
func (s *UserSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, CreateTableQuery)
}

// DropTable removes the users-table from the database
func (s *UserSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.Connection, DropTableQuery)
}
