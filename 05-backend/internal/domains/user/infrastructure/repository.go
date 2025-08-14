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

	// For batch operations, we still use the old approach since QueryBuilder is designed for single operations
	fields := []string{FieldEmail, FieldPassword, FieldDisplayName}
	setClauses := sql.CreateDollarClause(1, fields)
	whereClauses := sql.CreateDollarClause(4, []string{FieldId})
	query := sql.BuildUpdateQuery(TableName, setClauses, whereClauses)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetEmail(), entity.GetPassword(), entity.GetDisplayName(), entity.GetId()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Upsert inserts or updates one or more user entities in the database
func (s *UserRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	// For batch operations, we still use the old approach since QueryBuilder is designed for single operations
	insertFields := []string{FieldId, FieldEmail, FieldPassword, FieldDisplayName}
	keyFields := []string{FieldId}
	updateFields := []string{FieldEmail, FieldPassword, FieldDisplayName}
	query := sql.BuildUpsertQuery(TableName, insertFields, keyFields, updateFields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetEmail(), entity.GetPassword(), entity.GetDisplayName()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Delete removes one or more user entities from the database
func (s *UserRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UserEntity) error {
	if len(entities) == 0 {
		return nil
	}

	// For DELETE operations, we still use the old approach since QueryBuilder doesn't implement DELETE yet
	inClause := sql.CreateTupleInClause([]string{FieldId}, len(entities), 1)
	query := sql.BuildDeleteQuery(TableName, []string{inClause})

	values := make([]any, len(entities))
	for i, entity := range entities {
		values[i] = entity.GetId()
	}

	return database.Execute(ctx, s.Connection, query, values...)
}

// ReadOne executes a query and returns a single user entity
func (s *UserRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) (*domain.UserEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple user entities
func (s *UserRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserEntity]) ([]*domain.UserEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
