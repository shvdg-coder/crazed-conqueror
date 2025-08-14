package infrastructure

import (
	"context"
	"errors"
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
	query, args := sql.NewQuery().
		Select(FieldCharacterId, FieldFormationId).
		From(TableName).
		Where(FieldCharacterId, characterId).
		Build()
	return r.ReadMany(ctx, query, args, ScanCharacterFormationEntity)
}

// GetByFormationId retrieves all character formation entities for a given formation id
func (r *CharacterFormationRepositoryImpl) GetByFormationId(ctx context.Context, formationId string) ([]*domain.CharacterFormationEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldCharacterId, FieldFormationId).
		From(TableName).
		Where(FieldFormationId, formationId).
		Build()
	return r.ReadMany(ctx, query, args, ScanCharacterFormationEntity)
}

// Create inserts one or more character formation entities into the database
func (r *CharacterFormationRepositoryImpl) Create(ctx context.Context, entities ...*domain.CharacterFormationEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetCharacterId(), entity.GetFormationId()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldCharacterId, FieldFormationId).
		BatchValues(argumentSets).
		BuildBatch()

	return database.Batch(ctx, r.Connection, query, batchArgs)
}

// Update implements Repository.Update
func (r *CharacterFormationRepositoryImpl) Update(ctx context.Context, entities ...*domain.CharacterFormationEntity) error {
	return errors.New("update operation not supported for character-formation associations")
}

// Upsert implements Repository.Upsert
func (r *CharacterFormationRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.CharacterFormationEntity) error {
	return errors.New("upsert operation not supported for character-formation associations")
}

// Delete implements Repository.Delete
func (r *CharacterFormationRepositoryImpl) Delete(ctx context.Context, entities ...*domain.CharacterFormationEntity) error {
	return errors.New("delete operation not supported for character-formation associations")
}

// ReadOne executes a query and returns a single character formation entity
func (r *CharacterFormationRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterFormationEntity]) (*domain.CharacterFormationEntity, error) {
	return database.QueryOne(ctx, r.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple character formation entities
func (r *CharacterFormationRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.CharacterFormationEntity]) ([]*domain.CharacterFormationEntity, error) {
	return database.QueryMany(ctx, r.Connection, query, values, scan)
}
