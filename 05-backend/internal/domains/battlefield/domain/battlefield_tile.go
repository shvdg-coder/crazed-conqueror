package domain

import "shvdg/crazed-conquerer/internal/shared/types"

// BattlefieldTile represents a tile on the battlefield.
type BattlefieldTile struct {
	Coordinates *types.Coordinates
	OccupantIds []string
}

// NewBattlefieldTile creates a new instance of BattlefieldTile.
func NewBattlefieldTile(coordinates *types.Coordinates, occupantIds []string) *BattlefieldTile {
	return &BattlefieldTile{
		coordinates,
		occupantIds,
	}
}
