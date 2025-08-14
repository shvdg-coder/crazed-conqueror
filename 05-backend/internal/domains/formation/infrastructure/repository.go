package infrastructure

import (
	"context"
	"errors"
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
		Select(FieldId, FieldTiles, FieldCreatedAt, FieldUpdatedAt).
		From(TableName).
		Where(FieldId, id).
		Build()
	return s.ReadOne(ctx, query, args, ScanFormationEntity)
}

// Create implements Repository.Create
func (s *FormationRepositoryImpl) Create(ctx context.Context, entities ...*domain.FormationEntity) error {
	return errors.New("operation not supported")
}

// Update implements Repository.Update
func (s *FormationRepositoryImpl) Update(ctx context.Context, entities ...*domain.FormationEntity) error {
	return errors.New("operation not supported")
}

// Upsert implements Repository.Upsert
func (s *FormationRepositoryImpl) Upsert(ctx context.Context, entities ...*domain.FormationEntity) error {
	return errors.New("operation not supported")
}

// Delete implements Repository.Delete
func (s *FormationRepositoryImpl) Delete(ctx context.Context, entities ...*domain.FormationEntity) error {
	return errors.New("operation not supported")
}

// ReadOne executes a query and returns a single formation entity
func (s *FormationRepositoryImpl) ReadOne(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.FormationEntity]) (*domain.FormationEntity, error) {
	return database.QueryOne(ctx, s.Connection, query, values, scan)
}

// ReadMany executes a query and returns multiple formation entities
func (s *FormationRepositoryImpl) ReadMany(ctx context.Context, query string, values []any, scan database.ScannerFunc[*domain.FormationEntity]) ([]*domain.FormationEntity, error) {
	return database.QueryMany(ctx, s.Connection, query, values, scan)
}
