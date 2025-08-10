package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/unit/domain"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"shvdg/crazed-conquerer/internal/shared/database"

	"github.com/jackc/pgx/v5/pgtype"
)

// ScanUnitEntity scans database row data into a UnitEntity
func ScanUnitEntity(scanner database.RowScanner) (*domain.UnitEntity, error) {
	var unit domain.UnitEntity
	var createdAt, updatedAt pgtype.Timestamp

	err := scanner.Scan(
		&unit.Id,
		&unit.Vocation,
		&unit.Faction,
		&unit.Name,
		&unit.Level,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan unit entity: %w", err)
	}

	if createdAt.Valid {
		unit.CreatedAt = converters.TimeToTimestamp(createdAt.Time)
	}
	if updatedAt.Valid {
		unit.UpdatedAt = converters.TimeToTimestamp(updatedAt.Time)
	}

	return &unit, nil
}
