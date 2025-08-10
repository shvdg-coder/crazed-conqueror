package domain

import "context"

// CharacterRepository representation of a characterEntity repository
type CharacterRepository interface {
	GetById(ctx context.Context, id string) (*CharacterEntity, error)
}
