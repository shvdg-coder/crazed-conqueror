package domain

import "context"

// CharacterUnitRepository representation of a character unit repository
type CharacterUnitRepository interface {
	GetByCharacterID(ctx context.Context, characterID string) ([]*CharacterUnitEntity, error)
	GetByUnitID(ctx context.Context, unitID string) ([]*CharacterUnitEntity, error)
}
