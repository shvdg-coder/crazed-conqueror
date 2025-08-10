package infrastructure

import (
	"context"
	"errors"
	"shvdg/crazed-conquerer/internal/domains/character-unit/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// CharacterUnitRepositoryImpl provides the concrete implementation of the CharacterUnitRepository interface
type CharacterUnitRepositoryImpl struct {
	database.Connection
}

// NewCharacterUnitRepositoryImpl creates a new instance of CharacterUnitRepositoryImpl
func NewCharacterUnitRepositoryImpl(connection database.Connection) *CharacterUnitRepositoryImpl {
	return &CharacterUnitRepositoryImpl{connection}
}

// GetByCharacterId retrieves a character unit by their character ID
func (s *CharacterUnitRepositoryImpl) GetByCharacterId(ctx context.Context, characterID string) ([]*domain.CharacterUnitEntity, error) {
	fields := []string{FieldCharacterId, FieldUnitId}
	whereClause := sql.CreateDollarClause(1, []string{FieldCharacterId})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return s.ReadMany(ctx, query, []any{characterID}, ScanCharacterUnitEntity)
}

// GetByUnitId retrieves a character unit by their unit ID
func (s *CharacterUnitRepositoryImpl) GetByUnitId(ctx context.Context, unitID string) ([]*domain.CharacterUnitEntity, error) {
	fields := []string{FieldCharacterId, FieldUnitId}
	whereClause := sql.CreateDollarClause(1, []string{FieldUnitId})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return s.ReadMany(ctx, query, []any{unitID}, ScanCharacterUnitEntity)
}

// Create inserts one or more character unit entities into the database
func (s *CharacterUnitRepositoryImpl) Create(ctx context.Context, entities ...*domain.CharacterUnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldCharacterId, FieldUnitId}
	query := sql.BuildInsertQuery(TableName, fields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetCharacterId(), entity.GetUnitId()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Update is not supported for character-unit associations
func (s *CharacterUnitRepositoryImpl) Update(ctx context.Context, entities ...*domain.CharacterUnitEntity) error {
	return errors.New("update operation not supported for character-unit associations")
}

// Upsert is not supported for character-unit associations
func (s *CharacterUnitRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.CharacterUnitEntity) error {
	return errors.New("upsert operation not supported for character-unit associations")
}

// Delete removes one or more character unit entities from the database
func (s *CharacterUnitRepositoryImpl) Delete(ctx context.Context, entities ...*domain.CharacterUnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	inClause := sql.CreateTupleInClause([]string{FieldCharacterId, FieldUnitId}, len(entities), 1)
	query := sql.BuildDeleteQuery(TableName, []string{inClause})

	values := make([]any, len(entities)*2)
	for i, entity := range entities {
		values[i*2] = entity.GetCharacterId()
		values[i*2+1] = entity.GetUnitId()
	}

	return database.Execute(ctx, s.Connection, query, values...)
}

// ReadOne executes a query and returns a single character unit entity
func (s *CharacterUnitRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterUnitEntity]) (*domain.CharacterUnitEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple character unit entities
func (s *CharacterUnitRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterUnitEntity]) ([]*domain.CharacterUnitEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
