package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/character-formation/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// CharacterFormationRepositoryImpl provides the concrete implementation of the CharacterFormationRepository interface
type CharacterFormationRepositoryImpl struct {
	database.Connection
}

// NewCharacterFormationRepositoryImpl creates a new instance of CharacterFormationRepositoryImpl
func NewCharacterFormationRepositoryImpl(connection database.Connection) *CharacterFormationRepositoryImpl {
	return &CharacterFormationRepositoryImpl{connection}
}

// GetByCharacterId retrieves all character formation entities for a given character id
func (r *CharacterFormationRepositoryImpl) GetByCharacterId(ctx context.Context, characterId string) ([]*domain.CharacterFormationEntity, error) {
	fields := []string{FieldCharacterId, FieldFormationId}
	whereClause := sql.CreateDollarClause(1, []string{FieldCharacterId})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return r.ReadMany(ctx, query, []any{characterId}, ScanCharacterFormationEntity)
}

// GetByFormationId retrieves all character formation entities for a given formation id
func (r *CharacterFormationRepositoryImpl) GetByFormationId(ctx context.Context, formationId string) ([]*domain.CharacterFormationEntity, error) {
	fields := []string{FieldCharacterId, FieldFormationId}
	whereClause := sql.CreateDollarClause(1, []string{FieldFormationId})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return r.ReadMany(ctx, query, []any{formationId}, ScanCharacterFormationEntity)
}

// Create inserts one or more character formation entities into the database
func (r *CharacterFormationRepositoryImpl) Create(ctx context.Context, entities ...*domain.CharacterFormationEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldCharacterId, FieldFormationId}
	query := sql.BuildInsertQuery(TableName, fields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetCharacterId(), entity.GetFormationId()}
	}

	return database.Batch(ctx, r.Connection, query, argumentSets)
}

// ReadOne executes a query and returns a single character formation entity
func (r *CharacterFormationRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterFormationEntity]) (*domain.CharacterFormationEntity, error) {
	return database.QueryOne(ctx, r.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple character formation entities
func (r *CharacterFormationRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterFormationEntity]) ([]*domain.CharacterFormationEntity, error) {
	return database.QueryMany(ctx, r.Connection, query, values, scan)
}
