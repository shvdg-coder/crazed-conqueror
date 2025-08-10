package domain

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/types"
)

// Battlefield represents a battlefield on which units engage in combat.
type Battlefield struct {
	Tiles [][]BattlefieldTile
}

// GetByCoordinates retrieves a tile by row and column
func (b *Battlefield) GetByCoordinates(coordinates *types.Coordinates) (*BattlefieldTile, error) {
	if coordinates == nil {
		return nil, fmt.Errorf("coordinates cannot be nil")
	}

	if coordinates.GetX() >= int32(len(b.Tiles)) || coordinates.GetX() < 0 {
		return nil, fmt.Errorf("row %d out of bounds", coordinates.GetX())
	}
	if coordinates.GetY() >= int32(len(b.Tiles[coordinates.GetX()])) || coordinates.GetY() < 0 {
		return nil, fmt.Errorf("column %d out of bounds", coordinates.GetY())
	}

	return &b.Tiles[coordinates.GetX()][coordinates.GetY()], nil
}
