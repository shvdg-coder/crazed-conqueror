package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/user/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// UserRepositoryImpl provides the concrete implementation of the UserRepository interface
type UserRepositoryImpl struct {
	database.Connection
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(connection database.Connection) *UserRepositoryImpl {
	return &UserRepositoryImpl{connection}
}

// GetByEmail retrieves a user by their email address
func (s *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*domain.UserEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldId, FieldEmail, FieldPassword, FieldDisplayName, FieldCreatedAt, FieldUpdatedAt, FieldLastLoginAt).
		From(TableName).
		Where(FieldEmail, email).
		Build()
	return s.ReadOne(ctx, query, args, ScanUserEntity)
}

// Authenticate validates user credentials and returns the user if valid
func (s *UserRepositoryImpl) Authenticate(ctx context.Context, email, password string) (*domain.UserEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldId, FieldEmail, FieldPassword, FieldDisplayName, FieldCreatedAt, FieldUpdatedAt, FieldLastLoginAt).
		From(TableName).
		Where(FieldEmail, email).
		Where(FieldPassword, password).
		Build()
	return s.ReadOne(ctx, query, args, ScanUserEntity)
}

// Create inserts one or more user entities into the database
func (s *UserRepositoryImpl) Create(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldId, FieldEmail, FieldPassword, FieldDisplayName}
	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetEmail(), entity.GetPassword(), entity.GetDisplayName()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(fields...).
		BatchValues(argumentSets).
		BuildBatch()
	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Update modifies one or more user entities in the database
func (s *UserRepositoryImpl) Update(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argSets := make([][]any, len(entities))
	for i, entity := range entities {
		argSets[i] = []any{entity.GetEmail(), entity.GetPassword(), entity.GetDisplayName(), entity.GetId()}
	}

	query, batchArgs := sql.NewQuery().
		Update(TableName).
		BatchSets(argSets, FieldEmail, FieldPassword, FieldDisplayName).
		Where(FieldId).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Upsert inserts or updates one or more user entities in the database
func (s *UserRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argSets := make([][]any, len(entities))
	for i, entity := range entities {
		argSets[i] = []any{entity.GetId(), entity.GetEmail(), entity.GetPassword(), entity.GetDisplayName()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldId, FieldEmail, FieldPassword, FieldDisplayName).
		BatchUpsert(argSets, []string{FieldId}, FieldEmail, FieldPassword, FieldDisplayName).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Delete removes one or more user entities from the database
func (s *UserRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	ids := make([]any, len(entities))
	for i, entity := range entities {
		ids[i] = entity.GetId()
	}

	query, args := sql.NewQuery().
		DeleteFrom(TableName).
		WhereIn(FieldId, ids...).
		Build()

	return database.Execute(ctx, s.Connection, query, args...)
}

// ReadOne executes a query and returns a single user entity
func (s *UserRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) (*domain.UserEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple user entities
func (s *UserRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) ([]*domain.UserEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
