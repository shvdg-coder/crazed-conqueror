package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/user-character/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// ScanUserCharacterEntity scans database row data into a UserCharacterEntity
func ScanUserCharacterEntity(scanner database.RowScanner) (*domain.UserCharacterEntity, error) {
	var userCharacter domain.UserCharacterEntity

	err := scanner.Scan(
		&userCharacter.UserId,
		&userCharacter.CharacterId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan user character entity: %w", err)
	}

	return &userCharacter, nil
}
