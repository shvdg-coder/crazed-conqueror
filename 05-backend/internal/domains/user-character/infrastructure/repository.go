package infrastructure

import (
	"context"
	"errors"
	"shvdg/crazed-conquerer/internal/domains/user-character/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// UserCharacterRepositoryImpl provides the concrete implementation of the UserCharacterRepository interface
type UserCharacterRepositoryImpl struct {
	database.Connection
}

// NewUserCharacterRepositoryImpl creates a new instance of UserCharacterRepositoryImpl
func NewUserCharacterRepositoryImpl(connection database.Connection) *UserCharacterRepositoryImpl {
	return &UserCharacterRepositoryImpl{connection}
}

// GetByUserId retrieves all character associations for a given user ID
func (s *UserCharacterRepositoryImpl) GetByUserId(ctx context.Context, userID string) ([]*domain.UserCharacterEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldUserId, FieldCharacterId).
		From(TableName).
		Where(FieldUserId, userID).
		Build()
	return s.ReadMany(ctx, query, args, ScanUserCharacterEntity)
}

// GetByCharacterId retrieves all user associations for a given character ID
func (s *UserCharacterRepositoryImpl) GetByCharacterId(ctx context.Context, characterID string) ([]*domain.UserCharacterEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldUserId, FieldCharacterId).
		From(TableName).
		Where(FieldCharacterId, characterID).
		Build()
	return s.ReadMany(ctx, query, args, ScanUserCharacterEntity)
}

// Create stores user-character associations in the database
func (s *UserCharacterRepositoryImpl) Create(ctx context.Context, entities ...*domain.UserCharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldUserId, FieldCharacterId}
	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetUserId(), entity.GetCharacterId()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(fields...).
		BatchValues(argumentSets).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Update is not supported for user-character associations
func (s *UserCharacterRepositoryImpl) Update(ctx context.Context, entities ...*domain.UserCharacterEntity) error {
	return errors.New("update operation not supported for user-character associations")
}

// Upsert is not supported for user-character associations
func (s *UserCharacterRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.UserCharacterEntity) error {
	return errors.New("upsert operation not supported for user-character associations")
}

// Delete removes user-character associations from the database
func (s *UserCharacterRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UserCharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldUserId, FieldCharacterId}
	tuples := make([][]any, len(entities))
	for i, entity := range entities {
		tuples[i] = []any{entity.GetUserId(), entity.GetCharacterId()}
	}

	query, args := sql.NewQuery().
		DeleteFrom(TableName).
		WhereTupleIn(tuples, fields...).
		Build()

	return database.Execute(ctx, s.Connection, query, args...)
}

// ReadOne executes a query and returns a single user character entity
func (s *UserCharacterRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserCharacterEntity]) (*domain.UserCharacterEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple user character entities
func (s *UserCharacterRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserCharacterEntity]) ([]*domain.UserCharacterEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
