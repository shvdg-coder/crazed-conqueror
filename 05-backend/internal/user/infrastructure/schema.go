package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// UserSchema represents the user schema operations.
type UserSchema struct {
	connection database.Connection
}

// NewUserSchema creates a new instance of UserSchema.
func NewUserSchema(connection database.Connection) *UserSchema {
	return &UserSchema{
		connection: connection,
	}
}

// CreateTable creates the users-table in the database
func (s *UserSchema) CreateTable(ctx context.Context) error {
	return database.WithExecutor(ctx, s.connection, func(executor database.Executor) error {
		return database.Execute(ctx, executor, createTableQuery)
	})
}

// DropTable removes the users-table from the database
func (s *UserSchema) DropTable(ctx context.Context) error {
	return database.WithExecutor(ctx, s.connection, func(executor database.Executor) error {
		return database.Execute(ctx, executor, dropTableQuery)
	})
}
