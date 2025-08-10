package domain

import "context"

// CharacterFormationRepository representation of a character formation repository
type CharacterFormationRepository interface {
	GetByCharacterId(ctx context.Context, characterId string) ([]*CharacterFormationEntity, error)
	GetByFormationId(ctx context.Context, formationId string) ([]*CharacterFormationEntity, error)
}
