package domain

import "context"

// UnitRepository representation of a unit repository
type UnitRepository interface {
	GetById(ctx context.Context, id string) (*UnitEntity, error)
}
