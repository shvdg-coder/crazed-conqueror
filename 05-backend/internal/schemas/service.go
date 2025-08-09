package schemas

import (
	"context"
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/user/infrastructure"
)

// Service manages connection schemas across all domains.
type Service struct {
	connection database.Connection
	schemas    []database.DomainSchema
}

func NewDefaultService(connection database.Connection) *Service {
	return NewService(connection, infrastructure.NewUserSchema(connection))
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
