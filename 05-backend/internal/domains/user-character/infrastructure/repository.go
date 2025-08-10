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

// Create stores user-character associations in the database
func (s *UserCharacterRepositoryImpl) Create(ctx context.Context, entities ...*domain.UserCharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldUserID, FieldCharacterID}
	query := sql.BuildInsertQuery(TableName, fields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetUserId(), entity.GetCharacterId()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
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

	fields := []string{FieldUserID, FieldCharacterID}
	inClause := sql.CreateTupleInClause(fields, len(entities), 1)
	query := sql.BuildDeleteQuery(TableName, []string{inClause})

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetUserId(), entity.GetCharacterId()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// GetByUserID retrieves all character associations for a given user ID
func (s *UserCharacterRepositoryImpl) GetByUserID(ctx context.Context, userID string) ([]*domain.UserCharacterEntity, error) {
	fields := []string{FieldUserID, FieldCharacterID}
	whereClause := sql.CreateDollarClause(1, []string{FieldUserID})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return s.ReadMany(ctx, query, []any{userID}, ScanUserCharacterEntity)
}

// GetByCharacterID retrieves all user associations for a given character ID
func (s *UserCharacterRepositoryImpl) GetByCharacterID(ctx context.Context, characterID string) ([]*domain.UserCharacterEntity, error) {
	fields := []string{FieldUserID, FieldCharacterID}
	whereClause := sql.CreateDollarClause(1, []string{FieldCharacterID})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return s.ReadMany(ctx, query, []any{characterID}, ScanUserCharacterEntity)
}

// ReadOne executes a query and returns a single user character entity
func (s *UserCharacterRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserCharacterEntity]) (*domain.UserCharacterEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple user character entities
func (s *UserCharacterRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UserCharacterEntity]) ([]*domain.UserCharacterEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
