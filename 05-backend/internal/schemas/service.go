package schemas

import (
	"context"
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// Service manages connection schemas across all domains.
type Service struct {
	connection database.Connection
	schemas    []database.DomainSchema
}

// NewService creates a new schema service instance.
func NewService(db database.Connection, schemas ...database.DomainSchema) *Service {
	return &Service{
		connection: db,
		schemas:    schemas,
	}
}

// CreateAllTables creates all registered domain tables.
func (s *Service) CreateAllTables(ctx context.Context) error {
	for _, schema := range s.schemas {
		if err := schema.CreateTable(ctx); err != nil {
			return fmt.Errorf("failed to create tables: %w", err)
		}
	}

	return nil
}

// DropAllTables drops all registered domain tables in reverse order.
func (s *Service) DropAllTables(ctx context.Context) error {
	for i := len(s.schemas) - 1; i >= 0; i-- {
		if err := s.schemas[i].DropTable(ctx); err != nil {
			return fmt.Errorf("failed to drop tables: %w", err)
		}
	}

	return nil
}

// AddSchema adds a new domain schema to the service.
func (s *Service) AddSchema(schema database.DomainSchema) {
	s.schemas = append(s.schemas, schema)
}
