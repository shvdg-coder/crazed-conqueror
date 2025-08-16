package domain

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/types"
)

// Battlefield represents a battlefield on which units engage in combat.
type Battlefield struct {
	Rows []BattlefieldRow
}

// GetByCoordinates retrieves a column-by-row and column coordinates
func (b *Battlefield) GetByCoordinates(coordinates *types.Coordinates) (*BattlefieldColumn, error) {
	if coordinates == nil {
		return nil, fmt.Errorf("coordinates cannot be nil")
	}

	if coordinates.GetX() >= int32(len(b.Rows)) || coordinates.GetX() < 0 {
		return nil, fmt.Errorf("row %d out of bounds", coordinates.GetX())
	}
	if coordinates.GetY() >= int32(len(b.Rows[coordinates.GetX()].Columns)) || coordinates.GetY() < 0 {
		return nil, fmt.Errorf("column %d out of bounds", coordinates.GetY())
	}

	return &b.Rows[coordinates.GetX()].Columns[coordinates.GetY()], nil
}
