package domain

import "context"

// UserCharacterRepository representation of a user character repository
type UserCharacterRepository interface {
	GetByUserID(ctx context.Context, userID string) ([]*UserCharacterEntity, error)
	GetByCharacterID(ctx context.Context, characterID string) ([]*UserCharacterEntity, error)
}
