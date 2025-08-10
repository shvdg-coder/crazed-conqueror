package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/character/domain"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"shvdg/crazed-conquerer/internal/shared/database"

	"github.com/jackc/pgx/v5/pgtype"
)

// ScanCharacter scans database row data into a CharacterEntity entity
func ScanCharacter(scanner database.RowScanner) (*domain.CharacterEntity, error) {
	var character domain.CharacterEntity
	var createdAt, updatedAt pgtype.Timestamp

	err := scanner.Scan(
		&character.Id,
		&character.Name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan character entity: %w", err)
	}

	if createdAt.Valid {
		character.CreatedAt = converters.TimeToTimestamp(createdAt.Time)
	}
	if updatedAt.Valid {
		character.UpdatedAt = converters.TimeToTimestamp(updatedAt.Time)
	}

	return &character, nil
}
