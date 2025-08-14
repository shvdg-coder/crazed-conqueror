package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/character/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// CharacterRepositoryImpl provides the concrete implementation of the CharacterRepository interface
type CharacterRepositoryImpl struct {
	database.Connection
}

// NewCharacterRepositoryImpl creates a new instance of CharacterRepositoryImpl
func NewCharacterRepositoryImpl(connection database.Connection) *CharacterRepositoryImpl {
	return &CharacterRepositoryImpl{connection}
}

// GetById retrieves a character by their ID
func (s *CharacterRepositoryImpl) GetById(ctx context.Context, id string) (*domain.CharacterEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldId, FieldName, FieldCreatedAt, FieldUpdatedAt).
		From(TableName).
		Where(FieldId, id).
		Build()
	return s.ReadOne(ctx, query, args, ScanCharacter)
}

// Create inserts one or more character entities into the database
func (s *CharacterRepositoryImpl) Create(ctx context.Context, entities ...*domain.CharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetName()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldId, FieldName).
		BatchValues(argumentSets).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Update modifies one or more character entities in the database
func (s *CharacterRepositoryImpl) Update(ctx context.Context, entities ...*domain.CharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldName}
	setClauses := sql.CreateDollarClause(1, fields)
	whereClauses := sql.CreateDollarClause(2, []string{FieldId})
	query := sql.BuildUpdateQuery(TableName, setClauses, whereClauses)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetName(), entity.GetId()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Upsert inserts or updates one or more character entities in the database
func (s *CharacterRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.CharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetName()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldId, FieldName).
		BatchUpsert(argumentSets, []string{FieldId}, FieldName).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Delete removes one or more character entities from the database
func (s *CharacterRepositoryImpl) Delete(ctx context.Context, entities ...*domain.CharacterEntity) error {
	if len(entities) == 0 {
		return nil
	}

	inClause := sql.CreateTupleInClause([]string{FieldId}, len(entities), 1)
	query := sql.BuildDeleteQuery(TableName, []string{inClause})

	values := make([]any, len(entities))
	for i, entity := range entities {
		values[i] = entity.GetId()
	}

	return database.Execute(ctx, s.Connection, query, values...)
}

// ReadOne executes a query and returns a single character entity
func (s *CharacterRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterEntity]) (*domain.CharacterEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple character entities
func (s *CharacterRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterEntity]) ([]*domain.CharacterEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
