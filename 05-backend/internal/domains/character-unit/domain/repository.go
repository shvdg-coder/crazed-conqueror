package domain

import "context"

// CharacterUnitRepository representation of a character unit repository
type CharacterUnitRepository interface {
	GetByCharacterId(ctx context.Context, characterID string) ([]*CharacterUnitEntity, error)
	GetByUnitId(ctx context.Context, unitID string) ([]*CharacterUnitEntity, error)
}
