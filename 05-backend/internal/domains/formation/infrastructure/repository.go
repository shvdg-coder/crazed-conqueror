package infrastructure

import (
	"context"
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
	fields := []string{FieldId, FieldTiles, FieldCreatedAt, FieldUpdatedAt}
	whereClause := sql.CreateDollarClause(1, []string{FieldId})
	query := sql.BuildSelectQuery(TableName, fields, whereClause...)
	return s.ReadOne(ctx, query, []any{id}, ScanFormationEntity)
}

// ReadOne executes a query and returns a single formation entity
func (s *FormationRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.FormationEntity]) (*domain.FormationEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple formation entities
func (s *FormationRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.FormationEntity]) ([]*domain.FormationEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
