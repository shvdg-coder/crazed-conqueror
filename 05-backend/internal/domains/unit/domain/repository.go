package domain

import "context"

// UnitRepository representation of a unit repository
type UnitRepository interface {
	GetByID(ctx context.Context, id string) (*UnitEntity, error)
}
