package domain

import "context"

// UserCharacterRepository representation of a user character repository
type UserCharacterRepository interface {
	GetByUserId(ctx context.Context, userID string) ([]*UserCharacterEntity, error)
	GetByCharacterId(ctx context.Context, characterID string) ([]*UserCharacterEntity, error)
}
