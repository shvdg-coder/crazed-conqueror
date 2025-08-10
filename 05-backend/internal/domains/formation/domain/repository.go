package domain

import "context"

// FormationRepository representation of a formation repository
type FormationRepository interface {
	GetById(ctx context.Context, id string) (*FormationEntity, error)
}
