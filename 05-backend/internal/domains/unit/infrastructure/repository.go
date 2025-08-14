package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/unit/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
)

// UnitRepositoryImpl provides the concrete implementation of the UnitRepository interface
type UnitRepositoryImpl struct {
	database.Connection
}

// NewUnitRepositoryImpl creates a new instance of UnitRepositoryImpl
func NewUnitRepositoryImpl(connection database.Connection) *UnitRepositoryImpl {
	return &UnitRepositoryImpl{connection}
}

// GetById retrieves a unit by their ID
func (s *UnitRepositoryImpl) GetById(ctx context.Context, id string) (*domain.UnitEntity, error) {
	query, args := sql.NewQuery().
		Select(FieldId, FieldVocation, FieldFaction, FieldName, FieldLevel, FieldCreatedAt, FieldUpdatedAt).
		From(TableName).
		Where(FieldId, id).
		Build()
	return s.ReadOne(ctx, query, args, ScanUnitEntity)
}

// Create inserts one or more unit entities into the database
func (s *UnitRepositoryImpl) Create(ctx context.Context, entities ...*domain.UnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetVocation(), entity.GetFaction(), entity.GetName(), entity.GetLevel()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldId, FieldVocation, FieldFaction, FieldName, FieldLevel).
		BatchValues(argumentSets).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Update modifies one or more unit entities in the database
func (s *UnitRepositoryImpl) Update(ctx context.Context, entities ...*domain.UnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argSets := make([][]any, len(entities))
	for i, entity := range entities {
		argSets[i] = []any{entity.GetVocation(), entity.GetFaction(), entity.GetName(), entity.GetLevel(), entity.GetId()}
	}

	query, batchArgs := sql.NewQuery().
		Update(TableName).
		BatchSets(argSets, FieldVocation, FieldFaction, FieldName, FieldLevel).
		Where(FieldId).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Upsert inserts or updates one or more unit entities in the database
func (s *UnitRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.UnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	argSets := make([][]any, len(entities))
	for i, entity := range entities {
		argSets[i] = []any{entity.GetId(), entity.GetVocation(), entity.GetFaction(), entity.GetName(), entity.GetLevel()}
	}

	query, batchArgs := sql.NewQuery().
		InsertInto(TableName).
		InsertFields(FieldId, FieldVocation, FieldFaction, FieldName, FieldLevel).
		BatchUpsert(argSets, []string{FieldId}, FieldVocation, FieldFaction, FieldName, FieldLevel).
		BuildBatch()

	return database.Batch(ctx, s.Connection, query, batchArgs)
}

// Delete removes one or more unit entities from the database
func (s *UnitRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UnitEntity) error {
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

// ReadOne executes a query and returns a single unit entity
func (s *UnitRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UnitEntity]) (*domain.UnitEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple unit entities
func (s *UnitRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UnitEntity]) ([]*domain.UnitEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
