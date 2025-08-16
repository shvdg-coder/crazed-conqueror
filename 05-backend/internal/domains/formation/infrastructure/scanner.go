package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/formation/domain"
	"shvdg/crazed-conquerer/internal/shared/database"

	"github.com/jackc/pgx/v5/pgtype"
)

// ScanFormationEntity scans database row data into a FormationEntity
func ScanFormationEntity(scanner database.RowScanner) (*domain.FormationEntity, error) {
	var id string
	var rowsJson []byte
	var createdAt, updatedAt pgtype.Timestamp

	if err := scanner.Scan(&id, &rowsJson, &createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan formation entity: %w", err)
	}

	builder := domain.NewFormationEntity().
		WithId(id).
		WithRowsFromJson(rowsJson)

	if createdAt.Valid {
		builder = builder.WithCreatedAt(createdAt.Time)
	}
	if updatedAt.Valid {
		builder = builder.WithUpdatedAt(updatedAt.Time)
	}

	return builder.Build(), nil
}
