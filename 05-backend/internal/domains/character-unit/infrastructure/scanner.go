package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/character-unit/domain"
	"shvdg/crazed-conquerer/internal/shared/database"
)

// ScanCharacterUnitEntity scans database row data into a CharacterUnitEntity
func ScanCharacterUnitEntity(scanner database.RowScanner) (*domain.CharacterUnitEntity, error) {
	var characterUnit domain.CharacterUnitEntity

	err := scanner.Scan(
		&characterUnit.CharacterId,
		&characterUnit.UnitId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan character unit entity: %w", err)
	}

	return &characterUnit, nil
}
