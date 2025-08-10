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
	fields := []string{FieldId, FieldVocation, FieldFaction, FieldName, FieldLevel, FieldCreatedAt, FieldUpdatedAt}
	whereClause := sql.CreateDollarClause(1, []string{FieldId})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return s.ReadOne(ctx, query, []any{id}, ScanUnitEntity)
}

// Create inserts one or more unit entities into the database
func (s *UnitRepositoryImpl) Create(ctx context.Context, entities ...*domain.UnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldId, FieldVocation, FieldFaction, FieldName, FieldLevel}
	query := sql.BuildInsertQuery(TableName, fields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetVocation(), entity.GetFaction(), entity.GetName(), entity.GetLevel()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Update modifies one or more unit entities in the database
func (s *UnitRepositoryImpl) Update(ctx context.Context, entities ...*domain.UnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	fields := []string{FieldVocation, FieldFaction, FieldName, FieldLevel}
	setClauses := sql.CreateDollarClause(1, fields)
	whereClauses := sql.CreateDollarClause(5, []string{FieldId})
	query := sql.BuildUpdateQuery(TableName, setClauses, whereClauses)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetVocation(), entity.GetFaction(), entity.GetName(), entity.GetLevel(), entity.GetId()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Upsert inserts or updates one or more unit entities in the database
func (s *UnitRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.UnitEntity) error {
	if len(entities) == 0 {
		return nil
	}

	insertFields := []string{FieldId, FieldVocation, FieldFaction, FieldName, FieldLevel}
	keyFields := []string{FieldId}
	updateFields := []string{FieldVocation, FieldFaction, FieldName, FieldLevel}
	query := sql.BuildUpsertQuery(TableName, insertFields, keyFields, updateFields)

	argumentSets := make([][]any, len(entities))
	for i, entity := range entities {
		argumentSets[i] = []any{entity.GetId(), entity.GetVocation(), entity.GetFaction(), entity.GetName(), entity.GetLevel()}
	}

	return database.Batch(ctx, s.Connection, query, argumentSets)
}

// Delete removes one or more unit entities from the database
func (s *UnitRepositoryImpl) Delete(ctx context.Context, entities ...*domain.UnitEntity) error {
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

// ReadOne executes a query and returns a single unit entity
func (s *UnitRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UnitEntity]) (*domain.UnitEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple unit entities
func (s *UnitRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.UnitEntity]) ([]*domain.UnitEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
