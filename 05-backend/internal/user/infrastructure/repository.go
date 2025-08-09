package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
	"shvdg/crazed-conquerer/internal/user/domain"
)

// UserRepositoryImpl provides the concrete implementation of the UserRepository interface
type UserRepositoryImpl struct {
	executor database.Executor
	table    string
	fields   []string
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(executor database.Executor) domain.UserRepository {
	// TODO: Implementation
	return &UserRepositoryImpl{
		executor: executor,
		table:    "users",
		fields:   []string{"id", "email", "password", "display_name", "last_login_at", "created_at", "updated_at"},
	}
}

// GetByEmail retrieves a user by their email address
func (r *UserRepositoryImpl) GetByEmail(email string) (*domain.UserEntity, error) {
	// TODO: Implementation
	return nil, nil
}

// Authenticate validates user credentials and returns the user if valid
func (r *UserRepositoryImpl) Authenticate(email, password string) (*domain.UserEntity, error) {
	// TODO: Implementation
	return nil, nil
}

// Create inserts one or more user entities into the database
func (r *UserRepositoryImpl) Create(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	query := sql.BuildInsertQuery(r.table, r.fields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetEmail(), entity.GetPassword(),
			entity.GetDisplayName(), entity.GetLastLoginAt(),
			entity.GetCreatedAt(), entity.GetUpdatedAt(),
		}
	}

	return database.Batch(ctx, r.executor, query, argumentSets)
}

// Update modifies one or more user entities in the database
func (r *UserRepositoryImpl) Update(ctx context.Context, entities ...*domain.UserEntity) error {
	// TODO: Implementation
	return nil
}

// Delete removes one or more user entities from the database
func (r *UserRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UserEntity) error {
	// TODO: Implementation
	return nil
}

// ReadOne executes a query and returns a single user entity
func (r *UserRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) (*domain.UserEntity, error) {
	return nil, nil
}

// ReadMany executes a query and returns multiple user entities
func (r *UserRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) (*domain.UserEntity, error) {
	return nil, nil
}

// Count returns the number of user entities matching the given criteria
func (r *UserRepositoryImpl) Count(ctx context.Context, query string, values []any) (int, error) {
	// TODO: Implementation
	return 0, nil
}

// CreateTable creates the users-table in the database
func (r *UserRepositoryImpl) CreateTable(ctx context.Context) error {
	return database.Execute(ctx, r.executor, createTableQuery)
}

// DropTable removes the users-table from the database
func (r *UserRepositoryImpl) DropTable(ctx context.Context) error {
	return database.Execute(ctx, r.executor, dropTableQuery)
}
