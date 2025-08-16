package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/formation/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// FormationRepositoryImpl provides the concrete implementation of the FormationRepository interface
type FormationRepositoryImpl struct {
	database.Connection
}

// NewFormationRepositoryImpl creates a new instance of FormationRepositoryImpl
func NewFormationRepositoryImpl(connection database.Connection) *FormationRepositoryImpl {
	return &FormationRepositoryImpl{connection}
}

// GetById retrieves a formation by its id
func (s *FormationRepositoryImpl) GetById(ctx context.Context, id string) (*domain.FormationEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldId, FieldRows, FieldCreatedAt, FieldUpdatedAt).
		From(TableName).
		Where(FieldId, id).
		Build()

	return s.ReadOne(ctx, query, args, ScanFormationEntity)
}

// Create implements Repository.Create
func (s *FormationRepositoryImpl) Create(ctx context.Context, entities ...*domain.FormationEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argSets := make([][]any, len(entities))
	for i, entity := range entities {
		rows, err := json.Marshal(entity.GetRows())
		if err != nil {
			return fmt.Errorf("failed to marshal formation rows: %w", err)
		}

		argSets[i] = []any{entity.GetId(), json.RawMessage(rows)}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldId, FieldRows).
		BatchValues(argSets).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Update updates one or more formation entities in the database
func (s *FormationRepositoryImpl) Update(ctx context.Context, entities ...*domain.FormationEntity) error {
	return errors.New("operation not supported")
}

// Upsert upserts one or more formation entities in the database
func (s *FormationRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.FormationEntity) error {
	return errors.New("operation not supported")
}

// Delete removes one or more formation entities from the database
func (s *FormationRepositoryImpl) Delete(ctx context.Context, entities ...*domain.FormationEntity) error {
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

// ReadOne executes a query and returns a single formation entity
func (s *FormationRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.FormationEntity]) (*domain.FormationEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple formation entities
func (s *FormationRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.FormationEntity]) ([]*domain.FormationEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
