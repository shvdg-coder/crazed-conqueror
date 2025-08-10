package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/character-formation/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// ScanCharacterFormationEntity scans database row data into a CharacterFormationEntity
func ScanCharacterFormationEntity(scanner database.RowScanner) (*domain.CharacterFormationEntity, error) {
	var characterFormation domain.CharacterFormationEntity

	err := scanner.Scan(
		&characterFormation.CharacterId,
		&characterFormation.FormationId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan character formation entity: %w", err)
	}

	return &characterFormation, nil
}
