package domain

import "shvdg/crazed-conquerer/internal/shared/types"

// BattlefieldColumn represents a single column in the battlefield.
type BattlefieldColumn struct {
	Coordinates *types.Coordinates
	OccupantIds []string
}
