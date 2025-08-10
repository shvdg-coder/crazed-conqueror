package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
	"shvdg/crazed-conquerer/internal/user/domain"
)

// UserRepositoryImpl provides the concrete implementation of the UserRepository interface
type UserRepositoryImpl struct {
	connection database.Connection
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(connection database.Connection) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		connection: connection,
	}
}

// GetByEmail retrieves a user by their email address
func (s *UserRepositoryImpl) GetByEmail(email string) (*domain.UserEntity, error) {
	// TODO: Implementation
	return nil, nil
}

// Authenticate validates user credentials and returns the user if valid
func (s *UserRepositoryImpl) Authenticate(email, password string) (*domain.UserEntity, error) {
	// TODO: Implementation
	return nil, nil
}

// Create inserts one or more user entities into the database
func (s *UserRepositoryImpl) Create(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	return database.WithExecutor(ctx, s.connection, func(executor database.Executor) error {
		fields := []string{FieldId, FieldEmail, FieldPassword, FieldDisplayName}
		query := sql.BuildInsertQuery(TableName, fields)

		argumentSets := make([][]any, len(entities))
		for i, entity := range entities {
			argumentSets[i] = []any{entity.GetId(), entity.GetEmail(), entity.GetPassword(), entity.GetDisplayName()}
		}

		return database.Batch(ctx, executor, query, argumentSets)
	})
}

// Update modifies one or more user entities in the database
func (s *UserRepositoryImpl) Update(ctx context.Context, entities ...*domain.UserEntity) error {
	// TODO: Implementation
	return nil
}

// Delete removes one or more user entities from the database
func (s *UserRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UserEntity) error {
	// TODO: Implementation
	return nil
}

// ReadOne executes a query and returns a single user entity
func (s *UserRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) (*domain.UserEntity, error) {
	// TODO: Implementation
	return nil, nil
}

// ReadMany executes a query and returns multiple user entities
func (s *UserRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) (*domain.UserEntity, error) {
	// TODO: Implementation
	return nil, nil
}

// Count returns the number of user entities matching the given criteria
func (s *UserRepositoryImpl) Count(ctx context.Context, query string, values ...any) (int, error) {
	return database.WithExecutorResult(ctx, s.connection, func(executor database.Executor) (int, error) {
		var count int
		err := executor.QueryRow(ctx, query, values...).Scan(&count)
		if err != nil {
			return 0, err
		}
		return count, nil
	})
}
