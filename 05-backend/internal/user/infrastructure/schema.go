package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// UserSchema represents the user schema operations.
type UserSchema struct {
	executor database.Executor
}

// NewUserSchema creates a new instance of UserSchema.
func NewUserSchema(executor database.Executor) *UserSchema {
	return &UserSchema{
		executor: executor,
	}
}

// CreateTable creates the users-table in the database
func (s *UserSchema) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, s.executor, createTableQuery)
}

// DropTable removes the users-table from the database
func (s *UserSchema) DropTable(ctx context.Context) error {
	return database.Execute(ctx, s.executor, dropTableQuery)
}
